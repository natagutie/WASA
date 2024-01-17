<script>
export default {
	data: function() {
		return {
			errormsg: null,
      username: "",
      userID: "",
		}
	},
	methods: {

	async dologin(){
		try {
				
				console.log("Login method called")


				let response = await this.$axios.post("/session",{
                    username: this.username
                });
        
        console.log(response);
      
				localStorage.setItem("username", this.username);
        localStorage.setItem("userID", response.data);

        this.userID = response.data;
       
        // this.$axios.defaults.headers.common['Authorization']=`Bearer ${localStorage.getItem("userID")}`;




        this.$router.push({path: "/session"});
			} catch (e) {
				this.errormsg = e.toString();
			}
		
    },
	



	},
  
}
</script>

<template>
    
	  <h1 class="h1">Login</h1> <br><br>
    <div class="center">
  <input type="text" id="username" style="height: 40px" v-model="this.username" placeholder="Enter new username" ><br>
  <button class="btn btn-outline-dark" id="submit" @click="dologin" > submit  </button>
</div>
  </template>

<style>

.center {
  display: flex;
  align-items: center;
  justify-content: center;

}
.h1 {
  display: flex;
  flex-direction:row;
  align-items: center;
  justify-content: center;


}




</style>

