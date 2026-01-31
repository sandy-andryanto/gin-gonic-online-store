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
  selector: 'app-forgot-password-page',
  standalone: false,
  templateUrl: './forgot-password-page.component.html',
  styles: ``
})
export class ForgotPasswordPageComponent {

  forgotForm: FormGroup;
  nowYear: number = new Date().getFullYear()
  loading: boolean = false
  errorMessage:string = ""
  successMessage:string = ""
  private readonly router = inject(Router);

   constructor(private fb: FormBuilder, private authService: AuthService) {
    this.forgotForm = this.fb.group({
      email: ['', [Validators.required, Validators.email]],
    });
  }

  get email() {
    return this.forgotForm.get('email');
  }

  onSubmit() {

    if (this.forgotForm.invalid) {
      this.errorMessage = 'Form is invalid'
      return;
    }

    const { email } = this.forgotForm.value;

    this.loading = true
    this.successMessage = ""
    this.errorMessage = ""

    setTimeout(() => {
        this.authService.forgot({ email: email}).subscribe({
        next: (res) => {
          const token = res.token
          this.loading = false
          this.errorMessage = ""
          this.successMessage = res.message
          setTimeout(() => {
             this.router.navigate([`/auth/email/reset/${token}`]);
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
