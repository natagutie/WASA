<script>
export default {
  data: function() {
    return {
      userID: localStorage.getItem("userID"),
      username: localStorage.getItem("username"),
      errormsg: null,
      streamList: [
        {
          photoID: "",
          likeCount: "",
          comment: "",
          username: "",
          commentCount: "",
          date: "",
          photo: "",
          ifLike: false,
          text: "",
          fUsername: ""
        }
      ],
      commentList: [
        {
          commentID: 0,
          commentUser: "",
          comment: ""
        }
      ]
    };
  },
  created() {
    this.getMyStream();
  },
  methods: {
    async getMyStream() {
      try {
        let response = await this.$axios.get("/stream", {
          headers: {
            Authorization: "Bearer " + localStorage.getItem("userID")
          }
        });
        this.streamList = response.data;
        console.log(response.data);

        for (let i = 0; i < this.streamList.length; i++) {
          this.streamList[i].photo =
            "data:image/png;base64," + this.streamList[i].photoFile;
        }
      } catch (e) {
        this.errormsg = e.toString();
      }
    },
    async getUserComments(photoID) {
      let response = await this.$axios.get("/photo/" + photoID, {
        headers: {
          Authorization: "Bearer " + localStorage.getItem("userID")
        }
      });

      // Assuming commentList is an array
      this.commentList = response.data.map(comment => ({
        commentID: comment.commentID,
        photoID: comment.photoID,
        comment: comment.comment,
        username: comment.username
      }));

      var modal = document.getElementById("commentPopUp");
      modal.style.display = "block";
    },
    async RemoveComment(photoID, commentID) {
      await this.$axios.delete("/comment/" + commentID, {
        headers: {
          Authorization: "Bearer " + localStorage.getItem("userID")
        }
      });
      this.getUserComments(photoID);
      this.getMyStream();
    },
    async close() {
      var modal = document.getElementById("commentPopUp");
      modal.style.display = "none";
    },
    async LikePicture(photoID) {
      await this.$axios.post("/photo/" + photoID + "/like", {
        headers: {
          Authorization: "Bearer " + localStorage.getItem("userID")
        }
      });
      this.getMyStream();
    },
    async unLikePicture(photoID) {
      let response = await this.$axios.delete("/photo/" + photoID + "/like", {
        headers: {
          Authorization: "Bearer " + localStorage.getItem("userID")
        }
      });

      this.getMyStream();
    },
    async InsertComment(photoID, text) {
      await this.$axios.post(
        "/photo/" + photoID,
        { comment: text },
        {
          headers: {
            Authorization: "Bearer " + localStorage.getItem("userID")
          }
        }
      );
      this.getMyStream();
    }
  }
};
</script>

<template>
  <br /><br />
  <h2>Home</h2>
  <br />

  <div class="card-columns">
    <br /><br />
    <div class="card-columns">
      <div v-for="photo in this.streamList" :key="photo.photoID">
        <div class="card" style="width: 300px; height: 460px">
          <img
            class="card-img-top"
            :src="photo.photo"
            alt="error in showing photo"
            style="width=300px; height: 180px"
          />
          <h5><b> {{ photo.fUsername }}</b></h5>
          <div class="card-body">
            <button
              class="btn "
              v-if="photo.ifLike == true"
              @click="unLikePicture(photo.photoID)"
            >
              {{ photo.likeCount }}
            </button>
            <button
              class="btn"
              v-if="photo.ifLike == false"
              @click="LikePicture(photo.photoID)"
            >
              {{ photo.likeCount }}
            </button>
            <br /><button
              class="btn"
              @click="getUserComments(photo.photoID)"
            >
              Comments
            </button>

            <div
              class="modal "
              id="commentPopUp"
              tabindex="-1"
              role="dialog"
              aria-labelledby="exampleModalScrollableTitle"
            >
              <div
                class="modal-dialog modal-dialog-scrollable"
                role="document"
              >
                <div class="modal-content">
                  <span @click="close" class="close">&times;</span>
                  <div class="modal-header">
                    <h7 class="modal-title" style="color: #da91f1">Comments</h7>
                  </div>
                  <div class="modal-body">
                    <div
                      v-for="comment in this.commentList"
                      :key="comment.commentID"
                    >
                      <h6>
                        <b>{{ comment.username }}</b> {{ comment.comment }}
                      </h6>
                      <button
                        v-if="comment.username == this.username"
                        type="button"
                        class="btn"
                        style="float: right;"
                        @click="RemoveComment(
                          comment.photoID,
                          comment.commentID
                        )"
                      >
                        Delete
                      </button>
                      <br />
                    </div>
                  </div>
                </div>
              </div>
            </div>
            <p>
              <b>
                date: {{ photo.date }}<br />
                comments: {{ photo.commentCount }}<br />
              </b>{" "}
            </p>

            <div class="input-group mb-2">
              <input
                type="text"
                style="width: 200px"
                v-model="photo.text"
                class="form-control"
                placeholder="comment here!"
              />
              <div class="input-group-append">
                <button
                  class="btn"
                  style="height: 50px"
                  @click="InsertComment(photo.photoID, photo.text)"
                >
                  post
                </button>
              </div>
            </div>
            <br />
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
}

.card {
  border-color: pink;
}

.modal-body {
  background-color: pink;
}
</style>
