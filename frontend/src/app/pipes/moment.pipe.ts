/**
 * This file is part of the Sandy Andryanto Online Store Website.
 *
 * @author     Sandy Andryanto <sandy.andryanto.blade@gmail.com>
 * @copyright  2025
 *
 * For the full copyright and license information,
 * please view the LICENSE.md file that was distributed
 * with this source code.
 */

import { Pipe, PipeTransform } from '@angular/core';
import moment from 'moment';

@Pipe({
  name: 'moment',
  standalone: false
})
export class MomentPipe implements PipeTransform {
  transform(value: any, format: string = 'YYYY-MM-DD'): string {
    return moment(value).format(format);
  }
}
