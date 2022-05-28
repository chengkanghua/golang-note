<template>
  <el-container>
    <el-header>小清单
        <!-- <todoheader /> -->
        <button @click="logoutHandler" class="logout-btn">退出</button>
    </el-header>
    <el-main>
      <el-row type="flex" justify="center">
        <el-col :xs="20" :span="12" >
          <div class="grid-content">
            <router-view/>

          </div>
        </el-col>
      </el-row>
    </el-main>

    <el-footer> 友情链接：<el-link href="https://www.cnblogs.com/chengkanghua" target="_blank" type="primary">博客园</el-link> <el-link href="http://www.chengkanghua.top" target="_blank" type="primary">程康华的博客</el-link> <el-link href="https://www.liwenzhou.com/" target="_blank" type="primary">李文周的博客</el-link> 

    </el-footer>
  </el-container>
</template>
<script>
// import todoheader from "@/components/header.vue";
	// 导入组件
export default {
  name: 'App',
  // components:{
  //   // todoheader
  // },
  data(){
    return {

    }
  },
   methods:{
	  logoutHandler(){
		  this.$http.logout()
		  .then(res=>{
			  console.log(res)
			  if(res.code === 200){
				  this.$router.push({
					  name:"Login"
				  });
				  localStorage.removeItem('access_token')
				  localStorage.removeItem('name')
			  }
		  }).catch(err=>{
			  console.log(err);
		  })

	  },
  }
}
</script>

<style scoped>
.logout-btn{
  margin-top:15px;
  float: right;
}
.el-header,
.el-footer {
  background-color: #409eff;
  color: rgb(24, 24, 24);
  text-align: center;
  line-height: 60px;
}
.el-footer {
  background-color: #cfcfcf;
  display: block;
  width: 100%;
  position: fixed;
  bottom: 0;
}

</style>
