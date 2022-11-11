<template>
    <div class="login_box">
        <div class="inp">
            <el-input placeholder="用户名" v-model="name" clearable></el-input>
            <p> </p>
            <el-input placeholder="请输入密码" v-model="password" show-password></el-input>
            <el-button class="login_btn" type="primary" @click = 'loginHandler'>登录</el-button>
            
            <p class="go_register">没有账号 <el-button> <router-link :to = '{name:"Register"}'>立刻注册</router-link> </el-button></p>
        </div>
    </div>
</template>

<script>
export default {
    name: 'Login',
    data(){
        return {
            name:"",
            password:"",
        }
    },
    methods:{
        loginHandler(){
            let params = {
                name:this.name,
                password:this.password,
            }
            console.log(params)
            this.$http.userLogin(params)
            .then(res=>{
                console.log(res);
                this.$message(res.msg)
                if (res.code === 0){
                    this.$router.push({
                        name:"todolist"
                    });
                    localStorage.setItem('access_token',res.data);
                    localStorage.setItem('name',params.name)
                    
                    let recv ={
                        name:params.name,
                        access_token:res,
                    }
                    // dispacth action的行为
                    // this.$store.dispatch('getUserInfo',recv)
                }
            }).catch(err=>{
                console.log(err)
            })
        }
    }

}
</script>

<style scoped>

</style>