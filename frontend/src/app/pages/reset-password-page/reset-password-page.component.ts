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
import { Router, ActivatedRoute } from '@angular/router';


@Component({
  selector: 'app-reset-password-page',
  standalone: false,
  templateUrl: './reset-password-page.component.html',
  styles: ``
})
export class ResetPasswordPageComponent {

  nowYear: number = new Date().getFullYear()
  showPassword: boolean = false
  showPasswordConfirm: boolean = false
  loading: boolean = false
  resetForm: FormGroup;
  errorMessage:string = ""
  successMessage:string = ""

  private readonly router = inject(Router);

  constructor(private fb: FormBuilder, private authService: AuthService, private route: ActivatedRoute) {
    this.resetForm = this.fb.group({
      email: ['', [Validators.required, Validators.email]],
      password: ['', [Validators.required, Validators.minLength(6)]],
      passwordConfirm: ['', [Validators.required, Validators.minLength(6)]]
    });
  }

  setShowPassword() {
    this.showPassword = !this.showPassword
  }

  setShowPasswordConfirm() {
    this.showPasswordConfirm = !this.showPasswordConfirm
  }

  get email() {
    return this.resetForm.get('email');
  }

  get password() {
    return this.resetForm.get('password');
  }

  get passwordConfirm() {
    return this.resetForm.get('passwordConfirm');
  }

  onSubmit() {

    if (this.resetForm.invalid) {
      this.errorMessage = 'Form is invalid'
      return;
    }

    const { email, password, passwordConfirm } = this.resetForm.value;

    if(password !== passwordConfirm){
      this.errorMessage = 'Passwords must match'
      return;
    }

    this.loading = true
    this.errorMessage = ""
    this.successMessage = ""

    setTimeout(() => {
        const token = this.route.snapshot.paramMap.get('token') || '';
        this.authService.reset(token,{ email: email, password: password, password_confirm: passwordConfirm }).subscribe({
        next: (res) => {
          this.loading = false
          this.errorMessage = ""
          this.successMessage = res.message
          setTimeout(() => {
            this.router.navigate([`/auth/login`]);
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
