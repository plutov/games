class Games {
	constructor() {
		this.next('');
	}

	receive(onresult) {
		let self = this;

		let recognition = new webkitSpeechRecognition();
		recognition.onresult = onresult;
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
				if (r.status === 200) {
					self.say(r.responseText, function() {
						// TODO: check why is not always triggering
						self.tryToReceive();
					});
				} else {
					self.say('Server returns non-OK code ' + r.status);
				}
			}
		};
		r.open('GET', '/game?input=' + encodeURIComponent(userInput));
		r.send(null);
	}
}

new Games();
