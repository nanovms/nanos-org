const checker = document.querySelector('.navigate__mobile-menu');

checker.addEventListener('click', bodyChange)

function bodyChange(e) {
  if(e.target.checked) {
    window.onscroll = function () {
      window.scrollTo(0, 0)
    }
  } else {
    window.onscroll = function() {
      window.scrollTo(0, window.scrollY)
    }
  }
}
