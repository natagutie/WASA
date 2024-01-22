<script>
export default {
	data: function() {
		return {
			userID:"",
			userUsername:"",

			photos:[{
                photoID:0,
                username:"",
                like_count:0,
                comment_count:0,
                date:"",
                photo: "",
                isLiked:true,
                comment: "",
				followed_username:"",

                }
            ],
			    comments:[{
                    commentID:0,
                    commentUser: "",
                    comment:"",

                    }],
		}
	},
	created(){
		this.userID=localStorage.getItem("userID");

		this.getStream();
		
	},
	methods: {
		async getStream() {

			try {
				this.userID=localStorage.getItem("userID");
				this.userUsername= localStorage.getItem("username");

				let response = await this.$axios.get("/user/" + this.userID + "/stream", {
                headers: {
                    Authorization: "Bearer " + localStorage.getItem("userID")
                }
            });
				this.photos=response.data;


                for (let i=0;i<this.photos.length;i++){   
                    this.photos[i].photo= 'data:image/png;base64,'+ this.photos[i].photo;    
                }

			} catch (e) {
			}
			this.loading = false;
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
        
        
        this.getStream();
        this.getComments(photoID);
    
    },

	async LikePhoto(photoID){

        let response=await this.$axios.post("/user/" + this.userID + "/photo/" + photoID +"/like", {
                headers: {
                    Authorization: "Bearer " + localStorage.getItem("userID")
                }
        });
    this.getStream();
    
	},

	async unLikePhoto(photoID){
	let response=await this.$axios.delete("/user/" + this.userID + "/photo/" + photoID +"/like", {
                headers: {
                    Authorization: "Bearer " + localStorage.getItem("userID")
                }
            });

	this.getStream();
	},

	async uploadComment(photoID,yuhcomment){
           
           let response=await this.$axios.post("/user/" + this.userID + "/photo/" + photoID +"/comment", { comment: yuhcomment }, {
                headers: {
                    Authorization: "Bearer " + localStorage.getItem("userID")
                }
            });
           this.getStream();
	},

	async refresh() {
        this.getPhotos();
		this.getComments();
		
	},
	
},
	
}
</script>

<template>

	<h2 class="h22">My STREAM</h2>
			

<head>
<meta charset="utf-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
<link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap@4.6.2/dist/css/bootstrap.min.css">
<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/4.7.0/css/font-awesome.min.css"></head>
  
    

<div class="container">
	<div class="card-columns">
		<div v-for="photo in this.photos" style="width=300px" v-bind:key="photo.photoID">
			<div class="card" style="width: 300px; height: 500px">
				<RouterLink :to="'user/' + photo.followed_username + '/account'" class="nav-link">
				<img class="card-img-top" :src=photo.photo  alt="unavailable"  style="width=300px; height: 150px" >
			</RouterLink>
				<hr>
				<div class="card-body">
					<h6>{{ photo.followed_username }}</h6>
				<br>
				<button class="fa fa-heart-o" v-if="photo.isLiked==true" @click="this.unLikePhoto(photo.photoID)" :color="red"> {{ photo.like_count }}</button>
				<button class="fa fa-heart-o"  v-if="photo.isLiked==false" @click="this.LikePhoto(photo.photoID)" :color="green"> {{ photo.like_count }}</button>

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
                <button v-if="comment.username==this.userUsername" type="button" class="btn btn-danger" style="float: right;" @click="deleteComment(comment.comment_id, comment.photoID)">Delete</button>
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
                    </div><br>
				</div>
			
			</div>

		</div>
	</div>
	
	 
 </div>
 

	
</template>

<style>





.fa-heart-o {
  color: rgb(255, 255, 255);
  background-color: rgb(209, 161, 210);
  cursor: pointer;
  height: 27px;
}










</style>