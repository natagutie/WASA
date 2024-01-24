<script>
export default {
    data: function () {
        return {
            errormsg: null,
            ifPhotos: false,
            Username: localStorage.getItem("username"),
            newUsername: "",
            userID: localStorage.getItem("userID"),
            filePic: "",
            followersList: [],
            followingsList: [],
            bansList: [],
            comment:"",

            followerCount: null,
            followingCount: null,

            banCount: null,
            pictureCount: null,
            commentList: [{
                commentID: "",
                username: "",
                photoID:"",
                comment: "",

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

    created() {
        this.dataAccount();
    },

    methods: {

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
            



        async dataAccount() {

            try {
                this.Username = localStorage.getItem("username");
                this.userID = localStorage.getItem("userID");
                let response = await this.$axios.get("/profile/" + this.Username, {
                    headers: {
                        Authorization: "Bearer " + localStorage.getItem("userID")
                    }
                });

                this.banCount = response.data.banCount;
                this.bansList = response.data.bansList;
                this.followingCount = response.data.followingCount;
                this.pictureCount = response.data.photoCount;
                this.followerCount = response.data.followerCount;


                if (this.pictureCount > 0) {
                    this.getUserPhotos();
                    this.ifPhotos = true;
                };

            } catch (e) {
                this.errormsg = e.toString();
            }

        },


        async GetFollowingList() {
            try {
                let response = await this.$axios.get("/followings", {
                    headers: {
                        Authorization: "Bearer " + localStorage.getItem("userID")
                    }
                });
                this.followingsList = response.data;
                var modal = document.getElementById("followingPopUp");
                modal.style.display = "block";
            }
            catch (e) {
                this.errormsg = e.toString();
            }
        },



        close() {
            var modal = document.getElementById("commentPopUp");

            var followingPopUp = document.getElementById("followingPopUp");
            var banPopUp = document.getElementById("banPopUp");
            var followPopUp = document.getElementById("followPopUp");

            followingPopUp.style.display = "none";
            banPopUp.style.display = "none";
            followPopUp.style.display = "none";

            modal.style.display = "none";
        },
        async TheFollowers() {
        
                let response = await this.$axios.get("/followers", {
                    headers: {
                        Authorization: "Bearer " + localStorage.getItem("userID")
                    }
                });
                this.followersList = response.data;
                var modal = document.getElementById("followPopUp");
                modal.style.display = "block";
        
        },
        async GetBansList() {
            try {
                var modal = document.getElementById("banPopUp");
                modal.style.display = "block";
            }
            catch (e) {
                this.errormsg = "cant get ban List";
            }
        },


        async changeUsername() {
            if (this.newUsername == "") {
                this.errormsg = "The new username cant be empty!!!"
            } if (this.newUsername == this.username) {
                this.errormsg = "The new username cant be the same as your current"
            } else {
                try {
                    let response = await this.$axios.put("/username/" + this.newUsername, {
                        headers: {
                            Authorization: "Bearer " + this.userID
                        }
                    })
                    this.Username = response.data;
                    localStorage.setItem("username", this.Username);
                } catch (e) {
                    this.errormsg = "Error updating username";
                }
            }
        },
        async InsertPic() {
                try{     
                    console.log("pls work");
                    const aFile=this.$refs.picFile.files[0];
                    console.log(aFile);
                 
                    let response= await this.$axios.post("/photo", aFile, {
                        headers: {
                            Authorization: "Bearer " + localStorage.getItem("userID")
                        }
                    });

                    this.dataAccount();
                }
            catch(e){
                    this.errormsg=e.toString();
                }
        },
    


    async LikePicture(photoID) {
        try {
            await this.$axios.post("/photo/" + photoID + "/like", {
                headers: {
                    Authorization: "Bearer " + localStorage.getItem("userID")
                }
            });
            this.dataAccount();

        } catch (e) {
            this.errormsg = e.toString;


        }
    },

    async unLikePicture(photoID) {
        try {
            await this.$axios.delete("/photo/" + photoID + "/like", {
                headers: {
                    Authorization: "Bearer " + localStorage.getItem("userID")
                }
            });
            this.dataAccount();

        } catch (e) {
            this.errormsg = e.toString;
        }


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

    async getUserPhotos() {
        let pictures = await this.$axios.get("/photouser/" + this.Username, {
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
},
}


</script>
<template>
    <br>
    <h1> <b><span>{{ this.Username }}</span> 's Account </b></h1>

    <br>


    <div class="btn-group btn-group-sm">
				<input type="file" accept="image/*" class="btn btn-xs" id="inputImage" ref="picFile" @change="uploadFile" >
                <button class="btn btn-default " @click="InsertPic">Upload</button>
	</div>
   


    <div class="btn-group" style="float: right;">
    <button class="btn" @click="TheFollowers" > followers: {{ this.followerCount }} </button>

    <button class="btn" @click="GetFollowingList"> following: {{ this.followingCount }} </button>
    </div><br>
    <br>
    <div class="btn-group" style="float: right;">
    <button class="btn" @click="GetBansList" style="float: right;"> bans: {{ this.banCount }} </button>



    <button class="btn" style="float: right;"> photos: {{ this.pictureCount }} </button>
</div>

    <div class="modal" id="followPopUp" tabindex="-1" role="dialog">
        <div class="modal-dialog modal-dialog-scrollable" role="document">
            <div class="modal-content">
                <span @click="close" class="close">&times;</span>

                <div class="modal-header">
                    <h7 style="color: #da91f1" class="modal-title">Followers</h7>
                </div>
                <div class="modal-body">
                    <div v-for="follower in this.followersList" v-bind:key="follower">
                        <RouterLink :to="  '/account/'+ follower" class="nav-link">
                            <p> {{ follower }} </p>
                        </RouterLink>
                    </div>
                </div>
            </div>
        </div>
    </div>


    <div class="modal " id="followingPopUp" tabindex="-1" role="dialog">
        <div class="modal-dialog modal-dialog-scrollable" role="document">
            <div class="modal-content">
                <span @click="close" class="close">&times;</span>

                <div class="modal-header">
                    <h7 style="color: #da91f1" class="modal-title">Followings</h7>
                </div>

                <div class="modal-body">
                    <div v-for="following in this.followingsList" v-bind:key="following">
                        <RouterLink :to="  '/account/' + following" class="nav-link">
                            <p> {{ following }} </p>
                        </RouterLink>
                    </div>
                </div>
            </div>
        </div>
    </div>



    <div class="modal" id="banPopUp" tabindex="-1" role="dialog">
        <div class="modal-dialog modal-dialog-scrollable" role="document">
            <div class="modal-content">
                <span @click="close" class="close">&times;</span>

                <div class="modal-header">
                    <h7 style="color: #da91f1" class="modal-title"> Bans </h7>
                </div>
                <div class="modal-body">
                    <div v-for="baner in this.bansList" v-bind:key="baner.banID">
                        <RouterLink :to="'/account/' + baner" class="nav-link">
                            <p> {{ baner }} </p>
                        </RouterLink>
                    </div>
                </div>
            </div>
        </div>
    </div>


    <br>
    <br>
    <br>
    <div class="input-group mb-2">
    <input type="text" style="height: 40px;" id="username" v-model="this.newUsername" class="form-control" placeholder="Enter username"><br>
    <button class="btn" style="height: 40px" id="submit" @click="changeUsername"> submit </button>
    </div><br> <br> <br> <br> <br> <br> <br>


    <div class="container" v-if="this.pictureCount > 0">
        <div class="card-columns">
            <div v-for="photo in this.photoList" v-bind:key="photo.photoID">
                <div class="card" style="width: 300px; height: 400px">
                    <img class="card-img-top" :src=photo.photo alt="error in showing photo" style="width=300px; height: 200px">

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
                                likes: {{ photo.likesCount }}<br>
                                comments: {{ photo.commentCount }}<br>
                            </b> </p>

                            <div class="input-group mb-2">

                        <input type="text" style="width: 200px"  v-model="photo.text" class="form-control" placeholder="comment here!">
                        <div class="input-group-append">
                            <button class="btn" style="height: 50px" @click="InsertComment(photo.photoID, photo.text)">post</button>
                        </div>
                        </div><br>
                        <button class="btn " style="float: left;"
                            @click="deletePicture(photo.photoID)">Delete</button>
                    </div>
                </div>

            </div>
        </div>
    </div>
    <ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>

</template>


<style>
.btn {
    color: black;
    background-color: pink;
}

.card {
    border-color: pink;
}

.modal-body {
    background-color: pink;
}


</style>