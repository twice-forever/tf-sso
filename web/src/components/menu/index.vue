<template>
  <div class="mb-2">TF-SSO</div>
  <el-menu
    default-active="2"
    class="el-menu-vertical-demo"
    @open="handleOpen"
    @close="handleClose"
  >
    <el-menu-item index="1">
      <el-icon><bicycle /></el-icon>
      <span>Navigator One</span>
    </el-menu-item>
    <el-menu-item index="2">
      <el-icon><icon-menu /></el-icon>
      <span>Navigator Two</span>
    </el-menu-item>
    <el-menu-item index="3">
      <el-icon><document /></el-icon>
      <span>Navigator Three</span>
    </el-menu-item>
    <el-menu-item index="4">
      <el-icon><setting /></el-icon>
      <span>Navigator Four</span>
    </el-menu-item>
  </el-menu>
</template>

<script>
import {
  Bicycle,
  Menu as IconMenu,
  Document,
  Setting,
} from "@element-plus/icons-vue";

export default {
  name: "UserManage",
  components: {
    IconMenu,
    Setting,
    Bicycle,
    Document,
  },
  methods: {
    handleClose: (key, keyPath) => {
      console.log(key, keyPath);
    },
    handleOpen: (key, keyPath) => {
      console.log(key, keyPath);
    },
  },
};
</script>

<style scoped>
.mb-2 {
  text-align: center;
  padding: 16px 0;
  border: 1px solid #dcdfe6;
}
.el-menu {
  flex: 1;
}
</style>
