// src/vite-env.d.ts
/// <reference types="vite/client" />

// 声明 vite-plugin-svg-icons 生成的虚拟模块
declare module 'virtual:svg-icons-register' {
  const content: any;
  export default content;
}