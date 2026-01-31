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
export class ProfileService {

  constructor(private http: HttpClient) {}

  private authHeaders() : HttpHeaders {
    const token = localStorage.getItem('token');
    return new HttpHeaders({'Authorization': `Bearer ${token}`, 'Content-Type': 'application/json'});
  }

  detail(): Observable<any> {
    const headers = this.authHeaders()
    return this.http.get(`${environment.apiUrl}/api/profile/detail`, { headers });
  }

  update(data:any): Observable<any> {
    const headers = this.authHeaders()
    return this.http.post(`${environment.apiUrl}/api/profile/update`, data, { headers });
  }

  password(data:any): Observable<any> {
    const headers = this.authHeaders()
    return this.http.post(`${environment.apiUrl}/api/profile/password`, data, { headers });
  }

  upload(data:any): Observable<any> {
    const token = localStorage.getItem('token');
    const headers = new HttpHeaders({'Authorization': `Bearer ${token}`});
    return this.http.post(`${environment.apiUrl}/api/profile/upload`, data, { headers });
  }

  activity(params:string): Observable<any> {
    const headers = this.authHeaders()
    return this.http.get(`${environment.apiUrl}/api/profile/activity?${params}`, { headers });
  }

  getProfileImage(image:string){
    return `${environment.apiUrl}/uploads/${image}`
  }

}
