class Games {
	constructor() {
		this.terminal = document.getElementById('terminal');
		this.sessionID = this.terminal.getAttribute('data-session-id');
		this.next('');
	}

	receive(onresult) {
		let self = this;

		let recognition = new webkitSpeechRecognition();
		recognition.onresult = function(e) {
			if (e.results.length > 0 && e.results[0][0].transcript) {
				let p = document.createElement('p');
				p.innerHTML = '<em>YOU</em> > ' + e.results[0][0].transcript;
				self.terminal.appendChild(p);
			}
			onresult(e);
		}
		recognition.onerror = function(e) {
			self.say('An error has occurred with the speech recognition: ' + e.error);
		};
		recognition.start();
	}

	tryToReceive() {
		let self = this;

		self.receive(function(e) {
			if (e.results.length > 0 && e.results[0][0].transcript) {
				self.next(e.results[0][0].transcript);
			} else {
				self.say('Sorry, I can\'t help you with this, please try again.');
				self.tryToReceive();
			}
		});
	}

	say(text, onend) {
		if (!text) {
			onend();
			return;
		}

		// Add to terminal
		let p = document.createElement('p');
		p.innerHTML = '<em>ALEX</em> > ' + text;
		this.terminal.appendChild(p);

		let msg = new SpeechSynthesisUtterance(text);
		msg.onend = onend;
		msg.onerror = function(e) {
			console.log('An error has occurred with the speech synthesis: ' + e.error);
		};
		window.speechSynthesis.speak(msg);
	}

	next(userInput) {
		let self = this;
		let r = new XMLHttpRequest();

		r.onreadystatechange = function() {
			if (r.readyState === XMLHttpRequest.DONE) {
				if (r.status === 200 && r.response) {
					let json = JSON.parse(r.response);
					self.say(json.answer, function() {
						// TODO: check why is not always triggering
						if (json.game && !json.game.canceled) {
							self.tryToReceive();
						}
					});
				} else {
					self.say('Server returns non-OK code ' + r.status);
				}
			}
		};
		r.open('GET', '/game?input=' + encodeURIComponent(userInput) + '&session_id=' + this.sessionID);
		r.send(null);
	}
}

new Games();
