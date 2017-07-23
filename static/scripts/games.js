class Games {
	constructor() {
		let self = this;
		self.say('Hi, my name is Alex. Do you want to play a game?', function(e) {
			self.receive(function(e) {
				self.say('We have only one available game called "20 questions". Let\'s play?', function(e) {
					self.receive(function(e) {
						self.say('It\'s a simple and fun game where you should guess a noun by asking a maximum of 20 questions. Question can be answered only with "Yes", "No" and "Don\'t know. Let\'s play, try to ask something.', function(e) {
							self.receive(function(e) {
								self.say('Yes');
							});
						});
					});
				});
			});
		});
	}

	receive(onresult) {
		let self = this;
		let recognition = new webkitSpeechRecognition();
		recognition.onresult = onresult;
		recognition.start();
	}

	say(text, onend) {
		let msg = new SpeechSynthesisUtterance(text);
		msg.onend = onend;
		window.speechSynthesis.speak(msg);
	}
}

new Games();
