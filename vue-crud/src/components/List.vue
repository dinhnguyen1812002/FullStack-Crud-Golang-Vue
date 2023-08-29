<template>

    <div class="container">
        <h1>List All Book</h1>
        <hr>
        <table class="table table-hover">
            <tr>
                <th>STT</th>
                <th>Title</th>
                <th>Author</th>
                <th>Price</th>
                <th>Action</th>
            </tr>
            <tr v-for="(book, index) in books" :key="index">
                <td>{{ index + 1 }}</td>
                <td>{{ book.title }}</td>
                <td>{{ book.author }}</td>
                <td>{{ formatCurrencyVND(book.price) }}</td>
                <td>
                    <a @click="deleteBook(book)"  onclick="return confirm('Chac chua')" class="btn btn-danger">Delete</a>
                    <!-- <a @click="getID" class="btn btn-success">Update</a> -->
                    <router-link :to="'/book/'+ book.ID" class="btn btn-success">Edit</router-link>
                </td>
            </tr>
        </table>

    </div>

   
</template>
<script lang="ts">
import BookServices from "@/services/BookServices";
import ResponseData from "@/services/ResponseData";
import { Book } from "@/types/Book";
import { defineComponent } from "vue";
export default defineComponent({
    name: 'ListBook',
    data() {
        return {
            books: [] as Book[],
            currentBook: {} as Book,
            currentIndex: -1,

        };
    },
    methods: {
        getAll() {
            BookServices.getAll()
                .then((response: ResponseData) => {
                    this.books = response.data,
                        console.log(response.data)
                })
                .catch((e: Error) => {
                    console.log(e);
                });
        },
        formatCurrencyVND(value: number) {
            const formatter = new Intl.NumberFormat('vi-VN', {
                style: 'currency',
                currency: 'VND',
            });
            return formatter.format(value);
        },
        
        deleteBook(book: Book) {
            BookServices.delete(book.ID)
                .then((response: ResponseData) => {
                    console.log(response.data);
                    this.getAll()
                    this.$router.push({ name: "List" });
                })
                .catch((e: Error) => {
                    console.log(e);
                });
        }
    },
    mounted() {
        this.getAll()
    }
})
</script>
