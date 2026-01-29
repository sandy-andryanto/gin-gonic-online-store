/**
 * This file is part of the Sandy Andryanto Online Store Website.
 *
 * @author     Sandy Andryanto <sandy.andryanto.blade@gmail.com>
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
export class HomeService {

  constructor(private http: HttpClient) {}

  ping(): Observable<any> {
    return this.http.get(`${environment.apiUrl}/api/home/ping`);
  }

  component(): Observable<any> {
    return this.http.get(`${environment.apiUrl}/api/home/component`);
  }

  page(): Observable<any> {
    return this.http.get(`${environment.apiUrl}/api/home/page`);
  }

  newsletter(data:any): Observable<any> {
    return this.http.post(`${environment.apiUrl}/api/home/newsletter`, data);
  }

}
