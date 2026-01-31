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

import { Component, inject } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { AuthService } from '../../services/auth.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-register-page',
  standalone: false,
  templateUrl: './register-page.component.html',
  styles: ``
})
export class RegisterPageComponent {

  loading: boolean = false
  registerForm: FormGroup;
  nowYear: number = new Date().getFullYear()
  showPassword: boolean = false
  showPasswordConfirm: boolean = false
  errorMessage:string = ""

  private readonly router = inject(Router);

  constructor(private fb: FormBuilder, private authService: AuthService) {
    this.registerForm = this.fb.group({
      name: ['', [Validators.required]],
      email: ['', [Validators.required, Validators.email]],
      password: ['', [Validators.required, Validators.minLength(6)]],
      passwordConfirm: ['', [Validators.required, Validators.minLength(6)]]
    });
  }

  get name() {
    return this.registerForm.get('name');
  }

  get email() {
    return this.registerForm.get('email');
  }

  get password() {
    return this.registerForm.get('password');
  }

  get passwordConfirm() {
    return this.registerForm.get('passwordConfirm');
  }

  setShowPassword() {
    this.showPassword = !this.showPassword
  }

  setShowPasswordConfirm() {
    this.showPasswordConfirm = !this.showPasswordConfirm
  }

   onSubmit() {

    if (this.registerForm.invalid) {
      this.errorMessage = 'Form is invalid'
      return;
    }

    const { name, email, password, passwordConfirm } = this.registerForm.value;

    if(password !== passwordConfirm){
      this.errorMessage = 'Passwords must match'
      return;
    }

    this.loading = true
    this.errorMessage = ""
    setTimeout(() => {
        const formData = { name: name, email: email, password: password, password_confirm: passwordConfirm }
        this.authService.register(formData).subscribe({
        next: (res) => {
          const token = res.token
          this.loading = false
          this.errorMessage = ""
          setTimeout(() => {
            this.router.navigate([`/auth/register/confirm/${token}`]);
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
