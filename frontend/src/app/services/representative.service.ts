import {Injectable} from '@angular/core';
import {HttpClient} from "@angular/common/http";
import {Observable, shareReplay} from "rxjs";
import {Representative} from "../models/representative";

@Injectable({
  providedIn: 'root'
})
export class RepresentativeService {
  private BACKEND_URL = 'http://localhost:8080/api';

  constructor(private http: HttpClient) {
  }

  getMysqlRepresentatives(): Observable<Representative[]> {
    return this.http.get<Representative[]>(`${this.BACKEND_URL}/mysql/representatives`).pipe(shareReplay())
  }

  getPostgresRepresentatives(): Observable<Representative[]> {
    return this.http.get<Representative[]>(`${this.BACKEND_URL}/postgres/representatives`).pipe(shareReplay())
  }
}
