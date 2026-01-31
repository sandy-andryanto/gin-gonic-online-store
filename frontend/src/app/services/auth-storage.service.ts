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

@Injectable({
  providedIn: 'root'
})
export class AuthStorageService {

  constructor() { }

  saveToken(token: string) {
    localStorage.setItem('token', token);
  }

  saveUser(user: string) {
    const currentUser = this.getUser()
    if(currentUser === undefined || currentUser === null){
      localStorage.setItem('user', user);
    }
  }

  getUser(): string | null {
    return localStorage.getItem('user');
  }

  getToken(): string | null {
    return localStorage.getItem('token');
  }

  removeToken() {
    localStorage.removeItem('token');
  }

  removeUser() {
    localStorage.removeItem('user');
  }
}
