import apiClient from "@/http-commont";
class BookServices {
    getAll(): Promise<any> {
        return apiClient.get("/book");
    }
    delete(id: number): Promise<any>{
        return apiClient.delete(`/book/${id}`);
    }
    create(data: any):Promise<any>{
        return apiClient.post("/book",data)
    }
}
export default new BookServices;