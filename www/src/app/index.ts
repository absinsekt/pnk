import { isMobile } from './lib/detect';
import { initMasks } from './lib/masks';
import { initForms } from './lib/forms';

if (!isMobile()) {
  document.body.classList.add('desktop');
}

initMasks();
initForms();
