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

import { AfterViewInit, Component, inject } from '@angular/core';
import { AuthService } from '../../services/auth.service';
import { AnimationOptions } from 'ngx-lottie';
import { Router } from '@angular/router';
import { ActivatedRoute } from '@angular/router';

@Component({
  selector: 'app-confirm-page',
  standalone: false,
  templateUrl: './confirm-page.component.html',
  styles: ``
})
export class ConfirmPageComponent implements AfterViewInit {

  nowYear: number = new Date().getFullYear()
  loading: boolean = true
  errorMessage:string = ""
  checkingOption: AnimationOptions = { path: '/animations/checking.json' };
  confirmedOption: AnimationOptions = { path: '/animations/confirmed.json' };
  private readonly router = inject(Router);

  constructor(private authService: AuthService, private route: ActivatedRoute) {}

  ngAfterViewInit(): void {
    const token = this.route.snapshot.paramMap.get('token') || '';
    this.loading = true
    this.errorMessage = ""
    setTimeout(() => {
        this.authService.confirm(token).subscribe({
        next: (res) => {
          const message = res.message
          this.loading = false
          this.errorMessage = ""
        },
        error: (err) => {
          const message = err.error.error || err.error?.message || 'Something went wrong';
          this.errorMessage = message
          this.loading = false
        }
      });
    }, 2000)
  }


}
