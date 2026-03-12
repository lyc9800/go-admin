<template>
  <div class="login-container">
    <div class="login-box">
      <!-- 左侧视觉区域 -->
      <div class="login-left">
        <div class="image-container">
          <img src="@/assets/login/side-logo.png" alt="侧边背景" class="side-img" />
          <div class="image-overlay"></div>
        </div>
      </div>

      <!-- 登录表单区域 -->
      <div class="login-form">
        <!-- 品牌标识区 -->
        <div class="brand-area">
          <img class="logo" src="@/assets/login/logo02.svg" alt="后台管理系统" />
        </div>
        <!-- 登录表单组件 -->
        <LoginForm/>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import LoginForm from './components/LoginForm.vue'
</script>

<style scoped>
/* ========== 重置样式 ========== */
* {
  box-sizing: border-box;
  margin: 0;
  padding: 0;
}

/* ========== 背景容器 ========== */
.login-container {
  position: fixed;
  top: 0;
  left: 0;
  width: 100vw;
  height: 100vh;
  overflow: hidden;
  display: flex;
  justify-content: center;
  align-items: center;
  background: 
    linear-gradient(135deg, rgba(0, 0, 0, 0.02) 0%, rgba(0, 0, 0, 0) 100%),
    url("@/assets/login/login-bg.jpeg") center/cover no-repeat fixed;
  padding: 0;
  background-blend-mode: overlay;
  animation: bgFadeIn 1.2s ease-out;
}

@keyframes bgFadeIn {
  from {
    opacity: 0;
    background-position: center 10%;
  }
  to {
    opacity: 1;
    background-position: center center;
  }
}

/* ========== 登录主容器 ========== */
.login-box {
  position: relative;
  width: 94%;
  max-width: 1400px;
  height: 86%;
  min-height: 600px;
  max-height: 900px;
  background: linear-gradient(
    135deg,
    rgba(255, 255, 255, 0.98) 0%,
    rgba(255, 255, 255, 0.96) 50%,
    rgba(255, 255, 255, 0.94) 100%
  );
  border-radius: 24px;
  display: flex;
  align-items: stretch;
  justify-content: space-between;
  overflow: hidden;
  box-shadow: 
    0 20px 60px rgba(0, 0, 0, 0.15),
    0 8px 30px rgba(0, 0, 0, 0.08),
    0 3px 10px rgba(0, 0, 0, 0.05),
    inset 0 1px 0 rgba(255, 255, 255, 0.8),
    inset 0 -1px 0 rgba(255, 255, 255, 0.6);
  border: 1px solid rgba(255, 255, 255, 0.9);
  animation: containerFloat 0.8s cubic-bezier(0.16, 1, 0.3, 1) 0.2s both;
  position: relative;
  isolation: isolate;
}

@keyframes containerFloat {
  from {
    opacity: 0;
    transform: translateY(20px) scale(0.98);
  }
  to {
    opacity: 1;
    transform: translateY(0) scale(1);
  }
}

.login-box::before {
  content: '';
  position: absolute;
  top: -2px;
  left: -2px;
  right: -2px;
  bottom: -2px;
  background: linear-gradient(
    45deg,
    rgba(255, 255, 255, 0.1),
    rgba(255, 255, 255, 0.05),
    rgba(255, 255, 255, 0.1)
  );
  border-radius: 26px;
  z-index: -1;
  opacity: 0.6;
  filter: blur(8px);
}

/* ========== 左侧视觉区域 ========== */
.login-left {
  width: 80%;
  height: 100%;
  position: relative;
  overflow: hidden;
}

.image-container {
  position: relative;
  width: 100%;
  height: 100%;
  overflow: hidden;
}

.side-img {
  width: 100%;
  height: 100%;
  object-fit: cover;
  object-position: center;
  filter: blur(1.2px) brightness(0.97) saturate(1.1);
  transition: all 0.8s cubic-bezier(0.16, 1, 0.3, 1);
  transform-origin: center;
  transform: scale(1.02);
}

.login-left:hover .side-img {
  filter: blur(0.3px) brightness(1.05) saturate(1.2);
  transform: scale(1.05);
}

.image-overlay {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: linear-gradient(
    90deg,
    rgba(0, 0, 0, 0.15) 0%,
    rgba(0, 0, 0, 0.05) 20%,
    transparent 50%,
    transparent 100%
  );
  pointer-events: none;
  z-index: 1;
  opacity: 0.7;
  transition: opacity 0.6s ease;
}

.login-left:hover .image-overlay {
  opacity: 0.5;
}

/* ========== 登录表单区域（关键优化） ========== */
.login-form {
  width: 25%;
  padding: 40px 50px; /* 上下内边距40px，左右50px */
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center; /* 垂直居中！ */
  position: relative;
  animation: slideInRight 0.7s cubic-bezier(0.16, 1, 0.3, 1) 0.4s both;
  background: linear-gradient(
    to bottom,
    transparent 0%,
    rgba(255, 255, 255, 0.3) 50%,
    transparent 100%
  );
  background-size: 100% 200px;
  background-position: 0 -100px;
  background-repeat: no-repeat;
}

@keyframes slideInRight {
  from {
    opacity: 0;
    transform: translateX(40px);
  }
  to {
    opacity: 1;
    transform: translateX(0);
  }
}

/* ========== 品牌标识区 ========== */
.brand-area {
  display: flex;
  flex-direction: column;
  align-items: center;
  margin-bottom: -70px; /* 从40px改为20px，减少Logo与表单间距 */
  width: 100%;
  position: relative;
  animation: floatUp 0.8s cubic-bezier(0.16, 1, 0.3, 1) 0.6s both;
}

/* Logo样式 */
.login-form .logo {
  width: 320px; /* 稍微缩小一点，从260px改为240px */
  height: 220px; /* 从240px改为220px，更协调 */
  object-fit: contain;
  background-color: transparent;
  padding: 0;
  border: none;
  border-radius: 0;
  filter: 
    drop-shadow(0 8px 20px rgba(0, 0, 0, 0.12))
    drop-shadow(0 4px 8px rgba(245, 166, 35, 0.1))
    drop-shadow(0 2px 4px rgba(255, 255, 255, 0.3));
  transition: all 0.5s cubic-bezier(0.16, 1, 0.3, 1);
  cursor: pointer;
  transform-style: preserve-3d;
  transform: translateZ(0);
  -webkit-tap-highlight-color: transparent;
  user-select: none;
  /* 关键：调整向上移动距离，从-150px改为-50px */
  margin-top: -50px; /* 大幅减少上移距离 */
}
.login-form .logo:hover {
  transform: scale(1.08) translateY(-3px);
  filter: 
    drop-shadow(0 12px 30px rgba(0, 0, 0, 0.2))
    drop-shadow(0 6px 15px rgba(245, 166, 35, 0.18))
    drop-shadow(0 3px 8px rgba(255, 255, 255, 0.4))
    brightness(1.05);
}

.login-form .logo:active {
  transform: scale(1.04) translateY(-1px) !important;
  transition: transform 0.1s ease !important;
  filter: 
    drop-shadow(0 8px 20px rgba(0, 0, 0, 0.15))
    drop-shadow(0 4px 8px rgba(245, 166, 35, 0.12))
    drop-shadow(0 2px 4px rgba(255, 255, 255, 0.3)) !important;
}

/* ========== 登录表单组件样式 ========== */
:deep(.login-form-component) {
  width: 100%;
  max-width: 380px;
  margin-top: 0px; /* 从20px改为0px，让表单紧贴Logo */
  animation: fadeIn 0.7s ease 0.8s both;
  position: relative;
  z-index: 2;
  /* 控制表单内部的垂直间距 */
  --form-spacing: 16px; /* 定义变量方便调整 */
}

@keyframes fadeIn {
  from { 
    opacity: 0; 
    transform: translateY(10px);
  }
  to { 
    opacity: 1; 
    transform: translateY(0);
  }
}

/* ========== 性能优化 ========== */
.login-container,
.login-box,
.login-left,
.login-form {
  will-change: transform, opacity;
  backface-visibility: hidden;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}

.login-box,
.login-left,
.logo-container {
  transform: translateZ(0);
}

/* ========== 滚动条美化 ========== */
.login-form::-webkit-scrollbar {
  width: 6px;
}

.login-form::-webkit-scrollbar-track {
  background: rgba(0, 0, 0, 0.02);
  border-radius: 3px;
}

.login-form::-webkit-scrollbar-thumb {
  background: rgba(0, 0, 0, 0.1);
  border-radius: 3px;
}

.login-form::-webkit-scrollbar-thumb:hover {
  background: rgba(0, 0, 0, 0.15);
}
</style>