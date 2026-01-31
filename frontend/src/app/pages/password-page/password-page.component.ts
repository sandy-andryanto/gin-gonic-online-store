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
import { ProfileService } from '../../services/profile.service';

@Component({
  selector: 'app-password-page',
  standalone: false,
  templateUrl: './password-page.component.html',
  styles: ``
})
export class PasswordPageComponent {

  nowYear: number = new Date().getFullYear()
  showPasswordCurrent: boolean = false
  showPassword: boolean = false
  showPasswordConfirm: boolean = false
  loading: boolean = false
  passwordForm: FormGroup;
  errorMessage:string = ""
  successMessage:string = ""

  constructor(private fb: FormBuilder, private profileService: ProfileService) {
    this.passwordForm = this.fb.group({
      passwordCurrent: ['', [Validators.required, Validators.minLength(6)]],
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

  setShowPasswordCurrent() {
    this.showPasswordCurrent = !this.showPasswordCurrent
  }

  get password() {
    return this.passwordForm.get('password');
  }

  get passwordConfirm() {
    return this.passwordForm.get('passwordConfirm');
  }

  get passwordCurrent() {
    return this.passwordForm.get('passwordCurrent');
  }

  onSubmit() {

    if (this.passwordForm.invalid) {
      this.errorMessage = 'Form is invalid'
      return;
    }

    const { passwordCurrent, password, passwordConfirm } = this.passwordForm.value;

    if(password !== passwordConfirm){
      this.errorMessage = 'Passwords must match'
      return;
    }

    this.loading = true
    this.errorMessage = ""
    this.successMessage = ""

    setTimeout(() => {
        this.profileService.password({
          currentPassword: passwordCurrent,
          password: password,
          passwordConfirm: passwordConfirm
        }).subscribe({
        next: (res) => {
          this.loading = false
          this.errorMessage = ""
          this.successMessage = res.message
        },
        error: (err) => {
          const message = err.error?.message || 'Something went wrong';
          this.errorMessage = message
          this.loading = false
        }
      });
    }, 1500)

  }

}
