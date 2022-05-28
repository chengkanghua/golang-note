<template>
    <router-link :to = "{name:'Login'}" class="login_link" >登录</router-link>
</template>

<script>
// import TodoList from "@/components/TodoList.vue";
// import Login from "@/components/Login.vue";
export default {
  name: 'Header',
//   components:{
//     Login
//   }
}
</script>

<style scoped>
.login_link{
    float: right;
}

</style>
