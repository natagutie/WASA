<script>
export default {
	data: function() {
		return {
            errormsg: null,

            currentUserID: "",
            currentUser: "",
            
            userUsername: "",
            userIDUser:"",
            followerCount: 0,
            followingCount: 0,
            bannedCount:0,
            photoCount:0,
            isFollow: false,
            theyFollow: false,
            isBan: false,
            checkban:false,
            photos:[{
                photoID:0,
                username:"",
                like_count:0,
                comment_count:0,
                date:"",
                photo: "",
                isLike:false,
                comment: "",
                }],

            comments:[{
                commentUser: "",
                comment:"",
                    }],
		}
	},

    created(){
        this.CheckBanFollow()
    },

	methods: {

        async CheckBanFollow(){
            try {
                this.currentUserID=localStorage.getItem("userID");

                this.userUsername=this.$route.params.username;
                console.log(this.userUsername);

                this.currentUser=localStorage.getItem("username");
                let checkBan = await this.$axios.get("/username/"+ this.userUsername + "/ban/" + this.currentUser, {
                    headers: {
                        Authorization: "Bearer " + localStorage.getItem("userID")
                    }
                });
                this.checkban=checkBan.data;
                if(this.checkban){
                    this.$router.push({path: "/Error"});
                }else{
                let response = await this.$axios.get("/username/"+ this.currentUser + "/follow/" + this.userUsername, {
                    headers: {
                        Authorization: "Bearer " + localStorage.getItem("userID")
                    }
                });
   
                this.isFollow=response.data;
                if(this.isFollow){
                    document.getElementById("follow").innerHTML = "unfollow";
                    this.AccountInfo(); 
                }else{
                    document.getElementById("follow").innerHTML = "follow"; 
                    this.AccountInfo();
                }

                let banResponse = await this.$axios.get("/username/"+ this.currentUser + "/ban/" + this.userUsername, {
                    headers: {
                        Authorization: "Bearer " + localStorage.getItem("userID")
                    }
                });
 
                this.isBan=banResponse.data;
                if(this.isBan){
                    document.getElementById("ban").innerHTML = "unban"; 
                }else{
                    document.getElementById("ban").innerHTML = "ban"; 
                }
                this.refresh();
            }

                
            } catch (e) {
                this.errormsg = e.toString();
            }

        },

		 async AccountInfo() {
            
			try {
                let response=await this.$axios.get("/username/"+ this.userUsername + "/profile", {
                    headers: {
                        Authorization: "Bearer " + localStorage.getItem("userID")
                    }
                });
                    this.userIDUser=response.data.userID;
                    this.followerCount=response.data.followed_count;
                    this.followingCount=response.data.following_count;
                    this.photoCount=response.data.pic_numb;

               this.getPhotos();

               let response3 = await this.$axios.get("/username/"+ this.userUsername + "/follow/" + this.currentUser, {
                    headers: {
                        Authorization: "Bearer " + localStorage.getItem("userID")
                    }
                });
   
                this.theyFollow=response3.data;
                console.log("are they following?" + this.theyFollow);
            
			} catch (e) {
				this.errormsg = e.toString();
			}
		
    },
    async getComments(photoID) {
        let response=await this.$axios.get("/photo/" + photoID +"/comment", {
                headers: {
                    Authorization: "Bearer " + localStorage.getItem("userID")
                }
            });
        this.comments=response.data;
        var modal = document.getElementById("commentModal");
        modal.style.display = "block";
    },
    async closeModal() {
        var modal = document.getElementById("commentModal");

        modal.style.display = "none";
    },

    async deleteComment(commentID, photoID) {
        let response=await this.$axios.delete("/uncomment/"+commentID, {
                headers: {
                    Authorization: "Bearer " + localStorage.getItem("userID")
                }
            });
           
        this.refresh();
        this.getComments(photoID);

        
   
    
    },
          
    async getPhotos(){
           
           let response=await this.$axios.get("/user/" + this.currentUserID + "/photo/" + this.userIDUser, {
                headers: {
                    Authorization: "Bearer " + localStorage.getItem("userID")
                }
            });
           this.photos=response.data;

           for (let i=0;i<this.photos.length;i++){
               this.photos[i].photo= 'data:image/png;base64,'+ this.photos[i].photo;
           }
},

async uploadComment(photoID,yuhcomment){
      
    let response=await this.$axios.post("/user/" + this.userID + "/photo/" + photoID +"/comment", { comment: yuhcomment }, {
        headers:{
            Authorization: "Bearer " + localStorage.getItem("userID")
                }
    });
      this.AccountInfo();
      this.getPhotos();
},


    async getFollow(){
        try{     
            this.currentUser=localStorage.getItem("username");
            console.log(this.isFollow);
            if(this.isFollow){
                let response=await this.$axios.delete("/username/" + this.currentUser + "/follow/" + this.userUsername, {
                    headers: {
                        Authorization: "Bearer " + localStorage.getItem("userID")
                    }
                });
                this.isFollow=false;
                document.getElementById("follow").innerHTML = "follow"; 
                this.refresh();
            }else{
                console.log("ban check:" + this.isBan);
                if(this.isBan==false){
                let response=await this.$axios.post("/username/" + this.currentUser + "/follow/" + this.userUsername, {
                    headers: {
                        Authorization: "Bearer " + localStorage.getItem("userID")
                    }
                });
                this.isFollow=true;
                document.getElementById("follow").innerHTML = "unfollow";
                this.refresh();
                }
            
        }
       
    }
     catch(e){
            this.errormsg=e.toString();
        }


    },

    async refresh() {
		
        this.AccountInfo();
        
    },

    async getBan(){
        try{     
            this.currentUser=localStorage.getItem("username");
            if(this.isBan==true){
                let response=await this.$axios.delete("/username/" + this.currentUser + "/ban/" + this.userUsername, {
                    headers: {
                        Authorization: "Bearer " + localStorage.getItem("userID")
                    }
                });
                this.isBan=false;
                document.getElementById("ban").innerHTML = "ban"; 
                 this.refresh();
            }else{
                let response=await this.$axios.post("/username/" + this.currentUser + "/ban/" + this.userUsername, {
                    headers: {
                        Authorization: "Bearer " + localStorage.getItem("userID")
                    }
                });

                this.isBan=true;
                document.getElementById("ban").innerHTML = "unban";
                
                if(this.isFollow){
                    let response=await this.$axios.delete("/username/" + this.currentUser + "/follow/" + this.userUsername, {
                        headers: {
                            Authorization: "Bearer " + localStorage.getItem("userID")
                        }
                    });

                    this.isFollow=false;
                    document.getElementById("follow").innerHTML = "follow";                    
                }
                if(this.theyFollow){
                    let wtf=await this.$axios.delete("/username/" + this.userUsername + "/follow/" + this.currentUser, {
                            headers: {
                                Authorization: "Bearer " + localStorage.getItem("userID")
                            }
                        });
                    this.theyFollow=false;
                }
                   
        } 
        this.refresh(); 
    }
     catch(e){
        this.errormsg=e.toString();
        }

    },    
    async LikePhoto(photoID){

        console.log("putting like")

        let response=await this.$axios.post("/user/" + this.currentUserID + "/photo/" + photoID +"/like", {
                headers: {
                    Authorization: "Bearer " + localStorage.getItem("userID")
                }
            });
        this.getPhotos();
    },

    async unLikePhoto(photoID){

        console.log("deleting like")
        let response=await this.$axios.delete("/user/" + this.currentUserID + "/photo/" + photoID +"/like", {
                headers: {
                    Authorization: "Bearer " + localStorage.getItem("userID")
                }
        });
        console.log(response.data);

        this.getPhotos();
    },
    },
    }

 

</script>
<template>  

<head><meta name="viewport" content="width=device-width, initial-scale=1">
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/4.7.0/css/font-awesome.min.css"></head>
    
    <div class="titleButtons">

    <h2 > <span>{{this.userUsername}}</span> Account   </h2>

    </div>

    <div class="btn-group">
        <button class="btn btn-outline-dark " id="follow" @click="getFollow">wtv </button>

		<button class="btn btn-outline-dark " id="ban" @click="getBan">wtv</button>
	</div>
    
    <div class="btn-group">
 
            <button class="btn"> followers: {{this.followerCount}} </button>
    
            <button class="btn"> following: {{this.followingCount}} </button>
         
   
            <button class="btn"> photos: {{this.photoCount}} </button>
    </div>
    <br><br>
<head>
<meta charset="utf-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
<link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap@4.6.2/dist/css/bootstrap.min.css">
</head>

<div class="container">
    <div class="card-columns">
        <div v-for="photo in this.photos" style="width=300px" v-bind:key="photo.photoID">
            <div class="card" style="width: 300px; height: 500px">
                <img class="card-img-top" :src=photo.photo  alt="unavailable"  style="width=300px; height: 200px">
                <hr>
                <div class="card-body">
                    <h4>{{ photo.foll }}</h4>
                <button class="fa fa-heart-o"  @click="this.unLikePhoto(photo.photoID)" v-if="photo.isLiked==true"> {{ photo.like_count }}</button>
                <button class="fa fa-heart-o" @click="this.LikePhoto(photo.photoID)" v-if="photo.isLiked==false" > {{ photo.like_count }}</button>
                <button  @click="getComments(photo.photoID)"> comments</button>

            

<div class="modal " id="commentModal" tabindex="-1" role="dialog" aria-labelledby="exampleModalScrollableTitle" aria-hidden="true">
    <div class="modal-dialog modal-dialog-scrollable" role="document">
        <div class="modal-content">
            <span @click="closeModal" class="close">&times;</span>
        
            <div class="modal-header">
                <h5 class="modal-title">Comments</h5>
            </div>
            
            <div class="modal-body">
                <div v-for="comment in this.comments" v-bind:key="comment.commentID">
                <p> <b>{{ comment.username }}  </b> {{ comment.comment }} </p>
                <button v-if="comment.username==this.currentUser" type="button" class="btn btn-danger" style="float: right;" @click="deleteComment(comment.comment_id, photo.photoID)">Delete</button>
                <br>
                </div>
            </div>
        </div>
    </div>
</div>

                <br>
                <br>
                <p > likes: {{ photo.like_count }}<br>
                comments: {{ photo.comment_count }}<br>
                 date: {{ photo.date }}</p>

                <div class="gro">
                <input type="text" id="comment" v-model="photo.comment" class="form-control" placeholder="input comment">
                    <div class="input-group-append">
                        <button class="btn btn-outline-dark" type="button" @click="uploadComment(photo.photoID, photo.comment)">post</button>
                    </div>
                </div>
                </div>
            </div>
        </div>
    </div>
 </div>
 
	
</template>

<style>

.titleButtons{
    display:-webkit-flex;
    flex-direction:row;
    align-items:normal;
    justify-content:first baseline;
    margin-top: 37px;
    gap: 80px;

}


.btn-group{
  background-color: rgb(213, 213, 213);
  text-align: start;

}

.row{

    margin-top: 37px;
}

.card{
    box-shadow: 0 4px 8px 0 rgba(0,0,0,0.2);
    width: 100px;
    height:800px;
    

}


.container {
  margin-top: 10px;
  margin-right: 50px;
}


.gro{
    display:flex;
    flex-direction: row;
    align-items:flex-end;
    justify-content:start;
}

.fa-hearto {
  color: rgb(255, 255, 255);
  background-color: grey;
  cursor: pointer;
  height: 27px;
}
.fa-heart {
  color: grey;
  cursor: pointer;
  background-color: black;

  height: 27px;
}


</style>
