import {Injectable} from '@angular/core';
import {HttpClient} from "@angular/common/http";
import {Observable, shareReplay} from "rxjs";
import {CustomersListResponse} from "../models/customer";

@Injectable({
  providedIn: 'root'
})
export class CustomerService {
  private BACKEND_URL = 'http://localhost:8080/api';

  constructor(private http: HttpClient) {
  }

  getMysqlCustomers(params?: any): Observable<CustomersListResponse> {
    return this.http.get<CustomersListResponse>(`${this.BACKEND_URL}/mysql/customers`, {params: params}).pipe(shareReplay())
  }

  getPostgresCustomers(params?: any): Observable<CustomersListResponse> {
    return this.http.get<CustomersListResponse>(`${this.BACKEND_URL}/postgres/customers`, {params: params}).pipe(shareReplay())
  }
}
