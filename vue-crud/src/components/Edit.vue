<template>
    <div v-if="currentBook.ID">
        <h4>Edit book</h4>
        <div class="container">
            <form>
                <div class="submit-form">
                    <div class="row g-1">
                        <div class="col">
                            <!-- Name input -->
                            <div class="form-outline">
                                <input type="text" id="form9Example1" v-model="currentBook.title" class="form-control"
                                    placeholder="Title" />
                            </div>
                        </div>
                    </div>
                    <hr />
                    <!-- Gutter g-5 -->
                    <div class="row g-5">
                        <div class="col">
                            <!-- Name input -->
                            <div class="form-outline">
                                <input type="text" id="form9Example3" v-model="currentBook.author" class="form-control"
                                    placeholder="Author" />
                            </div>
                        </div>
                    </div>

                    <hr />
                    <div class="row g-1">
                        <div class="col">
                            <!-- Name input -->
                            <div class="form-outline">
                                <input type="number" id="form9Example1" v-model="currentBook.price" class="form-control"
                                    placeholder="price" />
                            </div>
                        </div>
                    </div>

                    <br>
                    <button @click="updateBook" class="btn btn-success">Update</button>
                </div>
                
            </form>

        </div>
    </div>
</template>
<script lang="ts">
import BookServices from "@/services/BookServices";
import ResponseData from "@/services/ResponseData";
import { Book } from "@/types/Book";
import { defineComponent } from "vue";
export default defineComponent({
    name: "EditForm",
    data() {
        return {
            currentBook: {} as Book,
            message: "",
        };
    },
    methods: {
        getBook(id: any) {
            BookServices.getByID(id)
                .then((response: ResponseData) => {
                    this.currentBook = response.data;
                    console.log(response.data);
                })
                .catch((e: Error) => {
                    console.log(e);
                });
        },
        updateBook() {
            BookServices.update(this.currentBook.ID, this.currentBook)
                .then((response: ResponseData) => {
                    console.log(response.data);
                    alert("The book was updated successfully!") 
                    this.$router.push({ name: "List" });
                })
                .catch((e: Error) => {
                    console.log(e);
                });
        },

    },
    mounted() {
        this.message = "";
        this.getBook(this.$route.params.ID);
    },

})
</script>