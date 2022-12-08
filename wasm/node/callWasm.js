require("./wasm_exec_tiny");
const fs = require("fs");

async function callWasm() {
  const go = new Go();

  const obj = await WebAssembly.instantiate(
    fs.readFileSync("./wasmbin"),
    go.importObject
  );
  const wasm = obj.instance;

  {
    // GoにStringを渡す
    const text = "Hello World!";
    const [addr, length] = writeBuffer(text, wasm);
    const result = wasm.exports.stringFromJS(addr, length);
    console.log(result); // Example: 文字数が帰ってくるように実装
    // Output: 12
  }

  {
    // GoからStringをもらう
    const addr = wasm.exports.stringToJS();
    const length = wasm.exports.getBufSize();
    const result = readBuffer(addr, length, wasm);
    console.log(result);
    // Output: Sample Str
  }
}

// 共有メモリを読み込む
function readBuffer(addr, size, module) {
  let memory = module.exports.memory;
  let bytes = memory.buffer.slice(addr, addr + size);
  let text = String.fromCharCode.apply(null, new Int8Array(bytes));
  return text;
}

// 共有メモリに書き込む
function writeBuffer(text, module) {
  // Get the address of the writable memory.
  const addr = module.exports.getBuffer();
  const buffer = module.exports.memory.buffer;

  const mem = new Int8Array(buffer);
  const view = mem.subarray(addr, addr + text.length);

  for (let i = 0; i < text.length; i++) {
    view[i] = text.charCodeAt(i);
  }

  // Return the address we started at.
  return [addr, text.length];
}

callWasm();

// 参考
// https://www.alcarney.me/blog/2020/passing-strings-between-tinygo-wasm/
// https://github.com/tinygo-org/tinygo/issues/3010
