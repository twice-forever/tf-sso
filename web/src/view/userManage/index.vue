<template>
  <div class="common-layout">
    <el-container>
      <el-aside width="200px"><Menu /></el-aside>
      <el-container>
        <el-header>Header</el-header>
        <el-main>
          
        </el-main>
        <el-footer>Footer</el-footer>
      </el-container>
    </el-container>
  </div>
</template>

<script>
import Menu from "@/components/menu/index.vue";

export default {
  name: "UserManage",
  components: {
    Menu,
  },
};
</script>

<style>
.common-layout .el-main {
  height: calc(100vh - 120px);
}
</style>
