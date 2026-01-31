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
import { HomeService } from '../../services/home.service';

@Component({
  selector: 'app-newsletter',
  standalone: false,
  templateUrl: './newsletter.component.html',
  styles: ``
})
export class NewsletterComponent {

  forgotForm: FormGroup;
  loading: boolean = false
  errorMessage:string = ""
  successMessage:string = ""

  constructor(private fb: FormBuilder, private homeService: HomeService) {
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
        this.homeService.newsletter({ email: email}).subscribe({
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
