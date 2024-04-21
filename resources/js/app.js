import '../css/styles.css';

import Alpine from 'alpinejs';

window.Alpine = Alpine;
Alpine.start();

document.addEventListener('DOMContentLoaded', () => {
});

document.addEventListener('alpine:init', () => {
    console.log('Alpine Loaded');
});
