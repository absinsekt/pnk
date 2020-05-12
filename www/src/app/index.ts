import { isMobile } from './core/detect';

if (!isMobile()) {
  document.body.classList.add('desktop');
}
