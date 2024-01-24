<script>
export default {
    components: {},
    data: function () {
        return {
            errormsg: null,
            username: "",
            userID:"",
 
        }
    },
    methods: {
        async logUser() {
            if (this.username == "") {
                this.errormsg = "Username cannot be empty.";
            } else {
                try {
                    let response = await this.$axios.post("/session", { username: this.username })
                    this.userID = response.data;
                    localStorage.setItem("userID", this.userID);
                    this.$axios.defaults.headers.common['Authorization']= Bearer ${localStorage.getItem("userID") };
                    localStorage.setItem("username", this.username);
                    this.$router.push({ path: '/session' })
                } catch (e) {
                    if (e.response) {
                        this.errormsg = "Error";
                        this.detailedmsg = null;
                    } 
                }
            }

        }
    },


}
</script>

<template>
  <br>
    <div>
        <h1>WASAPhoto</h1>
    </div>
    <br><br>
    <div class="input-group mb-2">

        <input type="text" style="width: 300px" id="username" v-model="this.username" class="form-control" placeholder="Username">
        <div class="input-group-append">
            <button class="btn" style="height: 40px" type="button" @click="logUser">Login</button>
        </div>
    </div>
    <ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
</template>

<style>
.btn {background-color: #e70d9e;}

</style>