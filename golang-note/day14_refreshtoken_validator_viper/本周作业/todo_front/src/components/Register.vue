<template>
    <div class="login_box">
        <div class="inp">
            <el-input placeholder="用户名" v-model="name" clearable></el-input>
            <p> </p>
            <el-input placeholder="请输入密码" v-model="password" show-password></el-input>
            <el-button class="login_btn" type="primary" @click = 'registerHandler'>注册</el-button>

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
        registerHandler(){
            let params = {
                name:this.name,
                password:this.password,
            }
            // console.log(params)
            this.$http.register(params)
            .then(res=>{
                console.log(res);
                if (res.code === 0){
                    this.$router.push({
                        name:"Login"
                    });
                    
                    // dispacth action的行为
                    // this.$store.dispatch('getUserInfo',recv)
                }else{
                    this.$message(res.msg)
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