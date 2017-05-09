document.addEventListener('DOMContentLoaded', event => {
    const music = document.getElementById('music');
    const playButton = document.getElementById('playButton');

    playButton.addEventListener('click', play);

    function play() {
        if (music.paused) {
            music.play();
            playButton.className = "playing";
        } else {
            music.pause();
            playButton.className = "paused";
        }
    };
});
