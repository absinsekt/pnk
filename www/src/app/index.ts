import { isMobile } from './core/detect';
import { initMasks } from './core/masks';
import { initForms } from './core/forms';

if (!isMobile()) {
  document.body.classList.add('desktop');
}

initMasks();
initForms();
