import Swal from 'sweetalert2';


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

  fetch(form.action, {
    method: form.method,
    mode: 'cors',
    cache: 'no-cache',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(serializeForm(form))
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
    .querySelectorAll('form[data-type="fetch"]')

    .forEach((form: HTMLFormElement) => {
      form.onsubmit = handleSubmit;
      form.onchange = handleChange;
    });
}

export { initForms }
