<script>
export default {
	data: function() {
		return {
			errormsg: null,
      searchUsername: "",
      userID: "",
      search: "",
      username:"",
      isban:true,
      isUser: false,
		}
	},
	methods: {

    async searchUser(){
		try {
            this.username=localStorage.getItem("username");
            this.userID=localStorage.getItem("userID");

            let response=await this.$axios.get("/username/"+ this.search + "/ban/" + this.username, {
                headers: {
                    Authorization: "Bearer " + localStorage.getItem("userID")
                }
            });
            this.isban=response.data
            console.log(this.isban);


            let responser=await this.$axios.get("/username/"+ this.search + "/checkUser", {
                headers: {
                    Authorization: "Bearer " + localStorage.getItem("userID")
                }
            });
            this.isUser=responser.data

            if(this.search==""){
                this.errormsg = "search cannot be empty"
            }else if(this.search==this.username){
                this.errormsg="Cant search yourself"

            }else if(this.isban==true){
              this.errormsg="Cant search this user."
            }else if(this.isUser==false){
              this.errormsg="username doesnt exist"
            }
            else{
              console.log(this.search)
                
              this.$router.push({path: "/account/" + this.search});
            }
        } 

        catch (e) {
				this.errormsg = e.toString();
		}
		
    },
	
    },
}
</script>

<template>
  		<br><br>
      <br><br>

		<div class="ok">
			<h2>Search User</h2>
		<br><br>
    </div>
    <br><br>

    <div class="input-group mb-2">
      <input type="text" style="width: 300px" id="username" v-model="this.search" class="form-control" placeholder="Username">
      <div class="input-group-append">
      <button class="btn" style="height: 40px" type="button" @click="searchUser">Search</button>
      </div>
    </div>
        <ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>


		
</template>
<style>


</style>