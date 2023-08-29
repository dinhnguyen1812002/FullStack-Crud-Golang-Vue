import apiClient from "@/http-commont";
class BookServices {
    getAll(): Promise<any> {
        return apiClient.get("/book");
    }
    getByID(id: number): Promise<any>{
        return apiClient.get(`/book/${id}`)
    }
    delete(id: number): Promise<any>{
        return apiClient.delete(`/book/${id}`);
    }
    create(data: any):Promise<any>{
        return apiClient.post("/book",data)
    }
    update(id:number,data:any){
        return apiClient.put(`/book/${id}`,data)
    }
}
export default new BookServices;