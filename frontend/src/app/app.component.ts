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

import { DOCUMENT } from '@angular/common';
import { Component, Inject, OnInit, Renderer2, HostListener } from '@angular/core';
import { AnimationOptions } from 'ngx-lottie';
import { HomeService } from './services/home.service';
import { AuthStorageService } from './services/auth-storage.service';


@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  standalone: false,
  styleUrl: './app.component.css'
})
export class AppComponent implements OnInit {

  title = '';
  loading: boolean = true
  connected: boolean = false
  backTo: boolean = false

  private document: any

  animationOption: AnimationOptions = { path: '/animations/loader.json' };

  constructor(
    @Inject(DOCUMENT) document: any,
    private render: Renderer2,
    private homeService: HomeService,
    private authStorageService: AuthStorageService
  ) {
    this.document = document
  }



  ngOnInit(): void {
      this.homeService.ping().subscribe({
        next: () => {
          setTimeout(async () => {
            this.loading = false
            this.connected = true
          }, 2000)
        },
        error: () => {
          setTimeout(()=> {
            this.loading = true
            this.connected = false
          })
        }
      });
  }

  clickToTop(event: any) {
    const e = event
    const window = this.document?.defaultView
    e.preventDefault();
    window.scrollTo({
      top: 0,
      behavior: 'smooth'
    });
    e.stopImmediatePropagation();
  }

  @HostListener('window:scroll', ['$event']) // for window scroll events
  onScroll(event: any) {
    this.backTo = this.document.body.scrollTop > 50 || this.document.documentElement.scrollTop > 50 ? true : false
  }


}
