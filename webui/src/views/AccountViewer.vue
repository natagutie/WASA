<script>
export default {
	data: function() {
		return {
            errormsg: null,
            currentUsername: localStorage.getItem("username"),
            userID: localStorage.getItem("userID"),
            ifPicture: false,


            profileUsername:"",

            bansList:[],
            followersList:[],
            followingsList:[],
 
            pictureCount:0,
            ifFollowing: false,
            ifBanned: false,

            commentList:[{
                        commentID: null,
                        username: "",
                        comment:"",

                    }],
        
            photoList: [{
                photoID: "",
                likeCount: "",
                comment: "",
                username: "",
                commentCount: "",
                date: "",
                photo: "",
                ifLiked: false,
                text: "",

            }
            ],
     
            }
		},
    created(){
    
        this.IfFollow();
        this.IfBanned();
        this.dataAccount();
        
    },

	methods: {
        async IfFollow(){
            try {
            this.profileUsername=this.$route.params.username;
            console.log(this.profileUsername);


            let response = await this.$axios.get("/follow/"+ this.profileUsername , {
                    headers: {
                        Authorization: "Bearer " + localStorage.getItem("userID")
                    }
            });
   
            this.ifFollowing=response.data;
            console.log(this.ifFollowing);
            if(this.ifFollowing==false){
                document.getElementById("follow").innerHTML = "follow"; 
            }else if(this.ifFollowing==false){
                document.getElementById("follow").innerHTML = "unfollow";

            }
            this.dataAccount();
            } catch (e) {
                this.errormsg = "failed getting users data";
            }

        },

    async getUserComments(photoID) {
            let response = await this.$axios.get("/photo/" + photoID, {
                headers: {
                    Authorization: "Bearer " + localStorage.getItem("userID")
                }
            });
            
            for (let i = 0; i < this.commentList.length; i++) {
            this.commentList[i].commentID = response.data[0].commentID;
            this.commentList[i].photoID = response.data[0].photoID;
            this.commentList[i].comment = response.data[0].comment;
            this.commentList[i].username = response.data[0].username;

            console.log(this.commentList[i].photoID);
            console.log(this.commentList[i].commentID);
        }

            var modal = document.getElementById("commentPopUp");
            modal.style.display = "block";
        },

        async RemoveComment(photoID, commentID) {
            await this.$axios.delete("/comment/" + commentID, {
                headers: {
                    Authorization: "Bearer " + localStorage.getItem("userID")
                }
            });

            this.dataAccount();
            this.getUserPhotos();
            this.getUserComments(photoID);
        },
    
    

    async close() {
        var modal = document.getElementById("commentPopUp");
        commentPopUp.style.display = "none";
        modal.style.display = "none";
    },

	async dataAccount() {
            
	    try {

            let response=await this.$axios.get("/profile/"+ this.profileUsername, {
                headers: {
                    Authorization: "Bearer " + localStorage.getItem("userID")
                }
            });
                this.banCount = response.data.banCount;
                this.bansList = response.data.bansList;
                this.pictureCount = response.data.photoCount;

            if(this.pictureCount>0){
                this.ifPicture=true;
                this.getUserPhotos();
            };
            
		} catch (e) {
				this.errormsg = e.toString();
			}
		
    },

     async IfBanned(){
            try {
                this.profileUsername=this.$route.params.username;
                
                let ifban = await this.$axios.get("/username/" +this.currentUsername + "/ban/" + this.profileUsername, {
                    headers: {
                        Authorization: "Bearer " + localStorage.getItem("userID")
                    }
                });
 
                this.ifBanned=ifban.data;
                if(this.ifBanned==true){
                    document.getElementById("ban").innerHTML = "unban"; 
                }else if(this.ifBanned==false){
                    document.getElementById("ban").innerHTML = "ban"; 
                }
                this.dataAccount();
              
            } catch (e) {
                this.errormsg = e.toString();
            }

        },
        
    async LikePicture(photoID) {
            try{
                await this.$axios.post("/photo/" + photoID + "/like", {
                    headers: {
                        Authorization: "Bearer " + localStorage.getItem("userID")
                    }
                });
                this.dataAccount();

            }catch(e){
                this.errormsg = e.toString;


            }
        },

    async unLikePicture(photoID) {
            try{
                await this.$axios.delete( "/photo/" + photoID + "/like", {
                    headers: {
                        Authorization: "Bearer " + localStorage.getItem("userID")
                    }
                });
                this.dataAccount();

            }catch(e){
                this.errormsg=e.toString;
            }


        },
    async unfollow(){
        try{     
                await this.$axios.delete("/follow/" + this.profileUsername, {
                    headers: {
                        Authorization: "Bearer " + localStorage.getItem("userID")
                    }
                });
                document.getElementById("follow").innerHTML = "follow"; 
                this.dataAccount();
        }
     catch(e){
            this.errormsg=e.toString();
        }
    },
    async doFollow(){
        try{     
            if(this.ifFollowing==false){
                this.follow();
                this.ifFollowing=true;
                
            } else{
                this.unfollow();
                this.ifFollowing=false;

            }   
        }
    
     catch(e){
            this.errormsg=e.toString();
        }
    },
    async follow(){
        try{     
            if(this.ifBanned==false){
                await this.$axios.post("/follow/" + this.profileUsername, {
                    headers: {
                        Authorization: "Bearer " + localStorage.getItem("userID")
                    }
                });
                document.getElementById("follow").innerHTML = "unfollow";
                this.dataAccount();
            }    
        }
    
     catch(e){
            this.errormsg=e.toString();
        }
    },

    async doBan(){
        try{     
            if(this.ifBanned==true){
                this.unban();
                this.ifBanned=false;

                
            }else{
                this.ban();
                this.ifBanned=true;

            }
    }
     catch(e){
        this.errormsg=e.toString();
        }

    },  

    async unban(){
        try{     
            await this.$axios.delete("/ban/" + this.profileUsername, {
                headers: {
                    Authorization: "Bearer " + localStorage.getItem("userID")
                }
            });
            document.getElementById("ban").innerHTML = "ban"; 
            this.dataAccount();
            }
     catch(e){
        this.errormsg=e.toString();
        }

    },    
    async ban(){
        try{      
            await this.$axios.post("/ban/" + this.profileUsername, {
                headers: {
                        Authorization: "Bearer " + localStorage.getItem("userID")
                    }
            });

            document.getElementById("ban").innerHTML = "unban";
                
            if(this.ifFollowing){
                await this.$axios.delete("/username/" + this.currentUser + "/follow/" + this.userUsername, {
                    headers: {
                        Authorization: "Bearer " + localStorage.getItem("userID")
                    }
                });

                this.ifFollowing=false;
                document.getElementById("follow").innerHTML = "follow";  

                await this.$axios.delete("/username/" + this.userUsername + "/follow/" + this.currentUser, {
                    headers: {
                        Authorization: "Bearer " + localStorage.getItem("userID")
                    }
                });                 
                }
        this.dataAccount(); 
        } 
        
    
     catch(e){
        this.errormsg=e.toString();
        }

    },    
 

    async getBansList(){
        try{    
            var modal = document.getElementById("banPopUp");
            modal.style.display = "block";
        }
     catch(e){
            this.errormsg=e.toString();
        }
    },
       
    
    async getUserPhotos() {
        let pictures = await this.$axios.get("/photouser/" + this.profileUsername, {
            headers: {
                Authorization: "Bearer " + localStorage.getItem("userID")
            }
        });
        console.log(pictures.data);
        this.photoList = pictures.data;
        console.log(this.photoList);
        var length = this.photoList.length;

        for (let i = 0; i < length; i++) {
            this.photoList[i].photo = 'data:image/png;base64,' + this.photoList[i].photoFile;
        }
        console.log(this.photoList);
    },

    async deletePicture(photoID) {
        await this.$axios.delete("/photo/" + photoID, {
            headers: {
                Authorization: "Bearer " + localStorage.getItem("userID")
            }
        });
        this.dataAccount();
    },
    async InsertComment(photoID, text){
        await this.$axios.post("/photo/" + photoID , {comment: text}, {
                headers: {
                    Authorization: "Bearer " + localStorage.getItem("userID")
                }
            });
        this.dataAccount();
      
},

},
}

</script>
<template>
    <br>
    <br>

    <h1 > <b><span>{{this.profileUsername}}</span> 's Account   </b></h1>

<br>	<br>
    <div class="btn-group">
        <button class="btn " id="follow" @click="doFollow"> follow  </button>

		<button class="btn" id="ban" @click="doBan">ban</button>
	</div>
    <br>  <br>
            <h5 style="float: left;"> photos: {{this.pictureCount}} </h5>

<br><br><br><br>
    <div v-if="this.pictureCount > 0">
        <div class="card-columns">
            <div v-for="photo in this.photoList" v-bind:key="photo.photoID">
                <div class="card" style="width: 300px; height: 430px">
                    <img class="card-img-top" :src=photo.photo alt="error in showing photo" style="width=300px; height: 180px">

                    <div class="card-body">
                        <button class="btn " v-if="photo.ifLiked == true" @click="this.unLikePicture(photo.photoID)"> {{ photo.likeCount }}</button>
                        <button class="btn" v-if="photo.ifLiked == false" @click="this.LikePicture(photo.photoID)"> {{ photo.likeCount }}</button>
                        <br><button class="commentbtn" @click="getUserComments(photo.photoID)"> comments</button>

                        <div class="modal " id="commentPopUp" tabindex="-1" role="dialog"
                            aria-labelledby="exampleModalScrollableTitle">
                            <div class="modal-dialog modal-dialog-scrollable" role="document">
                                <div class="modal-content">
                                    <span @click="close" class="close">&times;</span>
                                    <div class="modal-header">
                                        <h7 class="modal-title" style="color: #da91f1">Comments</h7>
                                    </div>
                                    <div class="modal-body">
                                        <div v-for="comment in this.commentList" v-bind:key="comment.commentID">
                                            <h6> <b>{{ comment.username }} </b> {{ comment.comment }} </h6>
                                            <button v-if="comment.username == this.Username" type="button" class="btn" style="float: right;" @click="RemoveComment( comment.photoID, comment.commentID)">Delete</button>
                                            <br>
                                        </div>
                                    </div>
                                </div>
                            </div>
                        </div>
                        <p> <b>
                                date: {{ photo.date }}<br>
                                likes: {{ photo.likeCount }}<br>
                                comments: {{ photo.commentCount }}<br>
                            </b> </p>

                            <div class="input-group mb-2">

                    <input type="text" style="width: 200px"  v-model="photo.text" class="form-control" placeholder="comment here!">
                    <div class="input-group-append">
                        <button class="btn" style="height: 50px" @click="InsertComment(photo.photoID, photo.text)">post</button>
                    </div>
                    </div>
                         <br>
                        
                    </div>
                </div>

            </div>
        </div>
    </div>
        
</template>


<style>


.btn {
  color: black;
  cursor: pointer;
  background-color: pink;
  height: 27px;
}
.commentbtn {
  color: black;
  cursor: pointer;
  background-color: #da91f1;
  height: 27px;
}

.commenter{
    display:flex;
    flex-direction: row;
    align-items:flex-end;
    justify-content:flex-start;
 
}
.card{

    border-color: pink;  
}
.modal-body {
    background-color: pink;
    width: 500px;
    height:1000px;
}

.container {
  margin-top: 10px;
  margin-right: 50px;
}



.combtn{
    border-color: pink;
    margin-right: 200px;
    width: 50px;
    height:30px;

}

.container {
  margin-top: -100px;
  margin-right: 50px;
}

.group{
    display:flex;
    flex-direction: row;
    align-items:flex-end;
    justify-content:start;
}

</style>