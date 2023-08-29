<template>
    <h1>thêm sản phẩm</h1>

        <div class="container">
            <div class="submit-form">
                <div v-if="!submitted">
                    <div class="row g-1">
                        <div class="col">
                            <!-- Name input -->
                            <div class="form-outline">
                                <input type="text" id="form9Example1" v-model="Book.title" class="form-control"
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
                                <input type="text" id="form9Example3" v-model="Book.author" class="form-control"
                                    placeholder="Author" />
                            </div>
                        </div>
                    </div>
                </div>
                <hr />
                <div class="row g-1">
                    <div class="col">
                        <!-- Name input -->
                        <div class="form-outline">
                            <input type="number" id="form9Example1" v-model="Book.price" class="form-control"
                                placeholder="price" />
                        </div>
                    </div>
                </div>

                <br>
                <button @click="saveBook" class="btn btn-success">Submit</button>
            </div>
        </div>
</template>
<script lang="ts">
import BookServices from '@/services/BookServices';
import { Book } from '@/types/Book';
import { defineComponent } from 'vue'
import ResponseData from '@/services/ResponseData';

export default defineComponent({
    name: "CreateForm",
    data() {
        return {
            Book: {
                title: "",
                author: "",
                price: 0
            } as Book,
            submitted: false,
        }
    },
    methods: {
        saveBook() {
            let data = {
                title: this.Book.title,
                author: this.Book.author,
                price: this.Book.price
            };
            BookServices.create(data)
                .then((response: ResponseData) => {
                    this.Book.ID = response.data.id;
                    console.log(response.data);
                    this.submitted = true;
                    this.$router.push({ name: "List" });
                })
                .catch((e: Error) => {
                    console.log(e);
                });

        },
        newTutorial() {
            this.submitted = false;
            this.Book = {} as Book;
        },

    }

})
</script>