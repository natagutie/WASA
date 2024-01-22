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


            let response2=await this.$axios.get("/username/"+ this.search + "/checkUser", {
                headers: {
                    Authorization: "Bearer " + localStorage.getItem("userID")
                }
            });
            this.isUser=response2.data

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
                
              this.$router.push({path: "user/" + this.search + "/account"});
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

      <div class="gro">
          <input class="text" type="text" style="height: 40px; width: 400px" v-model="search" placeholder="Search User">
          <button class="btn btn-outline-secondary" id="submit" style="height: 40px;" @click="searchUser" > search  </button>
      </div>
        <ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>


		
</template>
<style>


.ok {
  display: flex;
  flex-direction:row;
  align-items: center;
  justify-content: center;


}
.gro{
  display: flex;
  align-items: center;
  justify-content: center;
}

</style>