import Swal from 'sweetalert2';
import { isSet } from './objects';


function serializeForm(form: HTMLFormElement): {[key: string]: string|number|boolean} {
  const result = {};

  for(let idx = 0; idx < form.elements.length ; idx++) {
    const el = form.elements.item(idx) as HTMLInputElement;

    if (el.type !== 'submit') {
      const name = el.name;
      const value = el.value;

      result[name] = value;
    }
  }

  return result;
}

function handleSubmit(e) {
  const form = this;

  e.preventDefault();
  e.stopPropagation();

  const headers = {
    'Content-Type': 'application/json',
  };

  const data = serializeForm(form);

  const csrfToken = data['gorilla.csrf.Token'];
  if (isSet(csrfToken)) {
    headers['X-CSRF-Token'] = csrfToken;
    delete data['gorilla.csrf.Token'];
  }

  fetch(form.action, {
    method: form.method,
    mode: 'cors',
    cache: 'no-cache',
    headers,
    body: JSON.stringify(data)
  })

  .then((data) => data.json())
  .then(() => Swal.fire('Ololo!',))
  .catch(() => Swal.fire({
    title: 'Error!',
    icon: 'error'
  }));
}

function handleChange(e) {
  const form = this;
  console.log('TODO validators');
}

function initForms() {
  document
    .querySelectorAll('form[data-type="managed"]')

    .forEach((form: HTMLFormElement) => {
      form.onsubmit = handleSubmit;
      form.onchange = handleChange;
    });
}

export { initForms }
