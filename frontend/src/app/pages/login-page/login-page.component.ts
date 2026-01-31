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

import { Component } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { AuthService } from '../../services/auth.service';

@Component({
  selector: 'app-login-page',
  standalone: false,
  templateUrl: './login-page.component.html',
  styles: ``
})
export class LoginPageComponent {

  loginForm: FormGroup;
  nowYear: number = new Date().getFullYear()
  showPassword: boolean = false
  loading: boolean = false
  errorMessage:string = ""

  constructor(private fb: FormBuilder, private authService: AuthService) {
    this.loginForm = this.fb.group({
      email: ['', [Validators.required, Validators.email]],
      password: ['', [Validators.required, Validators.minLength(6)]]
    });
  }

  get email() {
    return this.loginForm.get('email');
  }

  get password() {
    return this.loginForm.get('password');
  }

  setShowPassword() {
    this.showPassword = !this.showPassword
  }

  onSubmit() {

    if (this.loginForm.invalid) {
      this.errorMessage = 'Form is invalid'
      return;
    }

    const { email, password } = this.loginForm.value;

    this.loading = true
    this.errorMessage = ""
    setTimeout(() => {
        this.authService.login({ email: email, password: password }).subscribe({
        next: (res) => {
          localStorage.setItem('token', res.token)
          this.loading = false
          this.errorMessage = ""
          setTimeout(() => {
            window.location.href = '/'
          }, 1500)
        },
        error: (err) => {
          const message = err.error.error || err.error?.message || 'Something went wrong';
          this.errorMessage = message
          this.loading = false
        }
      });
    }, 1500)

  }

}
