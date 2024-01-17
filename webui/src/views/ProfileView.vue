<script>
export default {
	data: function() {
		return {
            errormsg: null,
            hasposts: false,
            userUsername: localStorage.getItem("username"),
            newUser: "",
            userID: localStorage.getItem("userID"),
            followerCount: null,
            followingCount: null,
            followers:[],
            following:[],
            bans:[],
            bannedCount:null,
            photoCount:null,
            isFollow: false,
            isBan: false,
            isLoading: true,
            likeID: 0,

        
            photos:[{
                photoID:0,
                username:"",
                like_count:0,
                comment_count:0,
                date:"",
                photo: "",
                isLiked:true,
                comment: "",
                }
            ],
            comments:[{
                        comment_id: null,
                        username: "",
                        comment:"",

                    }],
            }

        
		},
     created(){

        this.AccountInfo();
        
    },

	methods: {

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
    async deleteComment(commentID, photoID) {
        let response=await this.$axios.delete("/uncomment/"+commentID, {
                headers: {
                    Authorization: "Bearer " + localStorage.getItem("userID")
                }
            });
        
        this.AccountInfo();
        this.getComments(photoID);
    
    },

    async closeModal() {
        var modal = document.getElementById("commentModal");
        var followmodal = document.getElementById("followerModal");
        var followingmodal = document.getElementById("followingModal");
        var banModal = document.getElementById("banModal");

        followmodal.style.display = "none";
        followingmodal.style.display = "none";
        banModal.style.display = "none";
        modal.style.display = "none";
    },

	async AccountInfo() {
            
	    try {
            this.userUsername=localStorage.getItem("username");
            this.userID=localStorage.getItem("userID");
            let response=await this.$axios.get("/username/"+ this.userUsername + "/profile", {
                headers: {
                    Authorization: "Bearer " + localStorage.getItem("userID")
                }
            });
                    this.followerCount=response.data.followed_count;
                    this.followingCount=response.data.following_count;
                    this.photoCount=response.data.pic_numb;
                    this.bannedCount=response.data.ban_count;
                    this.bans=response.data.bans;

            if(this.photoCount>0){
                    this.getPhotos();
                    this.hasposts=true;
            };
            
		} catch (e) {
				this.errormsg = e.toString();
			}
		
    },
    async refresh() {
		this.AccountInfo();    
	},


    async LikePhoto(photoID){

    let response=await this.$axios.post("/user/" + this.userID + "/photo/" + photoID +"/like", {
                headers: {
                    Authorization: "Bearer " + localStorage.getItem("userID")
                }
    });
    this.getPhotos();
    console.log(response.data);

    },

    async unLikePhoto(photoID){

    let response=await this.$axios.delete("/user/" + this.userID + "/photo/" + photoID +"/like", {
                headers: {
                    Authorization: "Bearer " + localStorage.getItem("userID")
                }
            });
    this.getPhotos();

    },


    async updateUsername(){

        if(this.newUser==""){
            this.errormsg="you must enter a username!!!"

        }else{
            let response=await this.$axios.put("/user/"+ this.userID + "/setusername/" + this.newUser, {
                headers: {
                    Authorization: "Bearer " + localStorage.getItem("userID")
                }
            });
            localStorage.setItem("username", response.data);
            this.refresh();
        }


    },
	
    async getFollower(){
        try{     
            let response=await this.$axios.get("/user/"+ this.userID + "/followers", {
                headers: {
                    Authorization: "Bearer " + localStorage.getItem("userID")
                }
            });
            this.followers=response.data;
            var modal = document.getElementById("followerModal");
            modal.style.display = "block"; 
        }
     catch(e){
            this.errormsg=e.toString();
        }
    },
    async getFollowing(){
        try{     
           
            let response=await this.$axios.get("/user/"+ this.userID + "/followings", {
                headers: {
                    Authorization: "Bearer " + localStorage.getItem("userID")
                }
            });
            this.following=response.data;
            var modal = document.getElementById("followingModal");
            modal.style.display = "block";
        }
     catch(e){
            this.errormsg=e.toString();
        }
    },
    async getBans(){
        try{    
            var modal = document.getElementById("banModal");
            modal.style.display = "block";
        }
     catch(e){
            this.errormsg=e.toString();
        }
    },

    async uploadPhoto(){
        try{     
            const uploadpic=this.$refs.pico.files[0];
            let response= await this.$axios.post("/photo/upload/" + this.userID, uploadpic, {
                headers: {
                    Authorization: "Bearer " + localStorage.getItem("userID")
                }
            });
            console.log(response.data);
            this.refresh();
            }
     catch(e){
            this.errormsg=e.toString();
        }
    },
    async deletePhoto(photoID) {
        let response= await this.$axios.delete("/user/" + this.userID + "/photo/" + photoID +"/remove", {
                headers: {
                    Authorization: "Bearer " + localStorage.getItem("userID")
                }
            });
        this.getPhotos();
       
    },
       
    async getPhotos(){
        let response=await this.$axios.get("/user/" + this.userID + "/photo/" + this.userID , {
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
                headers: {
                    Authorization: "Bearer " + localStorage.getItem("userID")
                }
            });
        this.AccountInfo();
        this.getPhotos();
},
},
}

</script>
<template>
    <head>
        <meta name="viewport" content="width=device-width, initial-scale=1">
        <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/4.7.0/css/font-awesome.min.css">
    </head>
  
    

    <h2 > <span>{{this.userUsername}}</span> Account   </h2>

  
    <div class="upload">
    <div class="btn-group btn-group-sm">
					<input type="file" accept="image/*" class="btn btn-xs" id="inputImage" ref="pico" @change="uploadFile" >

                    <button class="btn btn-default " @click="uploadPhoto">Upload</button>
				</div>
    </div>
   
    <div class="umidk">
    <div class="btn-group">
 
            <button class="btn" @click="getFollower"> followers: {{this.followerCount}} </button>
    
            <button class="btn" @click="getFollowing"> following: {{this.followingCount}} </button>
            <button class="btn" @click="getBans"> bans: {{this.bannedCount}} </button>

         
   
            <button class="btn"> photos: {{this.photoCount}} </button>
    </div>
    </div>

<div class="modal " id="followerModal" tabindex="-1" role="dialog">
    <div class="modal-dialog modal-dialog-scrollable" role="document">
        <div class="modal-content">
            <span @click="closeModal" class="close">&times;</span>
        
            <div class="modal-header">
                <h5 class="modal-title">Followers</h5>
            </div>
            <div class="modal-body" >
                <div v-for="follow in this.followers" v-bind:key="follow">
                <RouterLink :to="'user/' + follow + '/account'" class="nav-link">
                    <p> {{ follow }}  </p>
                </RouterLink>
            </div>
            </div>
        </div>
    </div>
</div>


<div class="modal " id="followingModal" tabindex="-1" role="dialog">
    <div class="modal-dialog modal-dialog-scrollable" role="document">
        <div class="modal-content">
            <span @click="closeModal" class="close">&times;</span>
        
            <div class="modal-header">
                <h5 class="modal-title">Followings</h5>
            </div>
            
            <div class="modal-body">
                <div v-for="follow in this.following" v-bind:key="follow">
                <RouterLink :to="'user/' + follow + '/account'" class="nav-link">
                    <p> {{ follow }}  </p>
                </RouterLink>
            </div>
        </div></div>
    </div>
</div>



<div class="modal" id="banModal" tabindex="-1" role="dialog" >
    <div class="modal-dialog modal-dialog-scrollable" role="document">
        <div class="modal-content">
            <span @click="closeModal" class="close">&times;</span>
        
            <div class="modal-header">
                <h5 class="modal-title"> Bans </h5>
            </div>
            <div class="modal-body" >
                <div  v-for="ban in this.bans" v-bind:key="ban.banID">
                <RouterLink :to="'user/' + ban + '/account'" class="nav-link">
                    <p> {{ ban }}  </p>
                </RouterLink>
            </div></div>
        </div>
    </div>
</div>


<br>
<div class="gro">
    <input type="text" style="height: 40px;" id="username" v-model="this.newUser" placeholder="Enter new username" ><br>
    <button class="btn btn-outline-dark" id="submit" @click="updateUsername"> submit  </button>
    </div><br>

    <head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap@4.6.2/dist/css/bootstrap.min.css">
 
    </head>

    <div class="container" v-if="this.hasposts==true">
        <div class="card-columns">
            <div v-for="photo in this.photos" style="width=300px" v-bind:key="photo.photoID">
                <div class="card" style="width: 300px; height: 500px">
                    <img class="card-img-top" :src=photo.photo  alt="unavailable" style="width=300px; height: 200px" >
                    <hr>
                    <div class="card-body">
                        <button type="button" class="btn btn-danger" style="float: right;" @click="deletePhoto(photo.photoID)">Delete</button>

                    <button class="fa fa-heart-o" v-if="photo.isLiked==true" @click="this.unLikePhoto(photo.photoID)"> {{ photo.like_count }}</button>
                    <button class="fa fa-heart-o"  v-if="photo.isLiked==false" @click="this.LikePhoto(photo.photoID)"> {{ photo.like_count }}</button>
                    <button  @click="getComments(photo.photoID)"> comments</button>



<div class="modal " id="commentModal" tabindex="-1" role="dialog">
    <div class="modal-dialog modal-dialog-scrollable" role="document">
        <div class="modal-content">
            <span @click="closeModal" class="close">&times;</span>
        
            <div class="modal-header">
                <h5 class="modal-title">Comments</h5>
            </div>
            
            <div class="modal-body">
                <div v-for="comment in this.comments" v-bind:key="comment.commentID">
                <p> <b>{{ comment.username }}  </b> {{ comment.comment }} </p>
                <button v-if="comment.username==this.userUsername" type="button" class="btn btn-danger" style="float: right;" @click="deleteComment(comment.comment_id, photo.photoID)">Delete</button>
                <br>
                </div>
            </div>
        </div>
    </div>
</div>
                    
                    <p >  likes: {{ photo.like_count }}<br>
                    comments: {{ photo.comment_count }}<br>
                     date: {{ photo.date }}</p>
   
                    <div class="gro">
                    <input type="text" id="comment" v-model="photo.comment" class="form-control" placeholder="input comment">
                        <div class="input-group-append">
                            <button class="btn btn-outline-dark" type="button" @click="uploadComment(photo.photoID, photo.comment)">post</button>
                        </div>
                    </div><br>
                    </div>
                    </div>
                
                </div>
</div>
</div>
        
</template>


<style>


.fa-hearto {
  color: red;
  background-color: grey;
  cursor: pointer;
  height: 27px;
}
.fa-heart {
  color: red;
  cursor: pointer;
  background-color: red;
  height: 27px;
}

.commenter{
    display:flex;
    flex-direction: row;
    align-items:flex-end;
    justify-content:flex-start;
 
}
.card{
    box-shadow: 0 4px 8px 0 rgba(0,0,0,0.2);
    width: 100px;
    height:800px;
    border-color: blue;  
}



.combtn{
    border-color: black;
    margin-right: 200px;
    width: 50px;
    height:30px;
    

}

.container {
  margin-top: -100px;
  margin-right: 50px;
}

.gro{
    display:flex;
    flex-direction: row;
    align-items:flex-end;
    justify-content:start;
}

</style>
