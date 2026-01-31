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

import { Component, Input } from '@angular/core';

@Component({
  selector: 'app-best-seller',
  standalone: false,
  templateUrl: './best-seller.component.html',
  styles: ``
})
export class BestSellerComponent {

  @Input() products: any[] = []

  slideConfig: any = {
    "slidesToShow": 1,
    "slidesToScroll": 1,
    "autoplay": true,
    "autoplaySpeed": 2000
  };

}
