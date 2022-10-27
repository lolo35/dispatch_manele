<template>
    <div class="absolute top-0 bottom-0 left-0 right-0 bg-black opacity-30 z-10"></div>
    <div class="absolute top-0 bottom-0 left-0 right-0 z-20">
        <div class="flex flex-col w-full h-full items-center justify-center">
            <span class="loader"></span>
        </div>
    </div>
</template>
<style scoped>
.loader {
    width: 48px;
    height: 48px;
    background: #fff;
    border-radius: 50%;
    position: relative;
    animation: skLinRotate 1s ease-in-out infinite alternate;
  }
  .loader:after{
    content:"";
    position: absolute;
    inset: 5px;
    border-radius: 50%;
    border: 5px solid transparent;
    border-top-color: #ff3d00;
  }
  @keyframes skLinRotate {
    95% ,100% { transform: rotate(840deg) }
  }
</style>