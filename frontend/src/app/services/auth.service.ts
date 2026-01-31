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
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { environment } from '../../environments/environment.development';


@Injectable({
  providedIn: 'root'
})
export class AuthService {

  constructor(private http: HttpClient) {}

  login(credentials: { email: string; password: string }): Observable<any> {
    return this.http.post(`${environment.apiUrl}/api/auth/login`, credentials);
  }

  register(credentials: {name: string;  email: string; password: string, password_confirm: string }): Observable<any> {
    return this.http.post(`${environment.apiUrl}/api/auth/register`, credentials);
  }

  confirm(token:string): Observable<any> {
    return this.http.get(`${environment.apiUrl}/api/auth/confirm/${token}`);
  }

  forgot(credentials: {email: string }): Observable<any> {
    return this.http.post(`${environment.apiUrl}/api/auth/email/forgot`, credentials);
  }

  reset(token:string, credentials: {email: string; password: string; password_confirm: string }): Observable<any> {
    return this.http.post(`${environment.apiUrl}/api/auth/email/reset/${token}`, credentials);
  }

}
