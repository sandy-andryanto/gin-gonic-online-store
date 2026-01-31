/**
 * This file is part of the Sandy Andryanto Online Store Website.
 *
 * @author     Sandy Andryanto <sandy.andryanto.official@gmail.com>
 * @copyright  2025
 *
 * For the full copyright and license information,
 * please view the LICENSE.md file that was distributed
 * with this source code.
 */

import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders  } from '@angular/common/http';
import { Observable } from 'rxjs';
import { environment } from '../../environments/environment.development';

@Injectable({
  providedIn: 'root'
})
export class OrderService {

   constructor(private http: HttpClient) {}

  private authHeaders() : HttpHeaders {
    const token = localStorage.getItem('token');
    return new HttpHeaders({'Authorization': `Bearer ${token}`, 'Content-Type': 'application/json'});
  }

  wishlist(id: number): Observable<any> {
    const headers = this.authHeaders()
    return this.http.get(`${environment.apiUrl}/api/order/wishlist/${id}`, { headers });
  }

  session(): Observable<any> {
    const headers = this.authHeaders()
    return this.http.get(`${environment.apiUrl}/api/order/session`, { headers });
  }

  cart(id:number): Observable<any> {
    const headers = this.authHeaders()
    return this.http.get(`${environment.apiUrl}/api/order/cart/${id}`, { headers });
  }

  listReview(id:number): Observable<any> {
    const headers = this.authHeaders()
    return this.http.get(`${environment.apiUrl}/api/order/review/${id}`, { headers });
  }

  createReview(id:number, data:any): Observable<any> {
    const headers = this.authHeaders()
    return this.http.post(`${environment.apiUrl}/api/order/review/${id}`, data, { headers });
  }

  createCart(id:number, data:any): Observable<any> {
    const headers = this.authHeaders()
    return this.http.post(`${environment.apiUrl}/api/order/create/cart/${id}`, data, { headers });
  }

   checkoutInitial(): Observable<any> {
    const headers = this.authHeaders()
    return this.http.get(`${environment.apiUrl}/api/order/checkout/initial`, { headers });
  }

  checkoutSubmit(data:any): Observable<any> {
    const headers = this.authHeaders()
    return this.http.post(`${environment.apiUrl}/api/order/checkout/submit`, data, { headers });
  }

  list(param:string): Observable<any> {
    const headers = this.authHeaders()
    return this.http.get(`${environment.apiUrl}/api/order/list${param}`, { headers });
  }

  detail(id:number): Observable<any> {
    const headers = this.authHeaders()
    return this.http.get(`${environment.apiUrl}/api/order/detail/${id}`, { headers });
  }

   cancel(id:number): Observable<any> {
    const headers = this.authHeaders()
    return this.http.get(`${environment.apiUrl}/api/order/cancel/${id}`, { headers });
  }

}
