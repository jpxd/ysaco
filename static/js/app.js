Number.prototype.pad = function(size) {
	var s = String(this);
	while (s.length < (size || 2)) {
		s = "0" + s;
	}
	return s;
}

function getCookie(name) {
    return (name = new RegExp('(?:^|;\\s*)' + ('' + name).replace(/[-[\]{}()*+?.,\\^$|#\s]/g, '\\$&') + '=([^;]*)').exec(document.cookie)) && name[1];
}

function playSample(id, start, duration, sender) {
	if (sender != null) {
		$(sender).addClass('loader');
	}
	var url = `//www.youtube.com/embed/${id}?start=${start}&end=${start + duration}&rel=0&autoplay=1`;
	$("#yt-video")[0].src = url;
}

function openModal() {
	clearModal();
	$('.modal-overlay').show();
	$('.modal').show();
}

function closeModal() {
	$('.modal').hide();
	$('.modal-overlay').hide();
}

function clearModal() {
	$('#name').val('');
	$('#url').val('');
	$('#from').val('');
	$('#duration').val('');
}

function clearLoaders() {
	setTimeout(function(){
		$('a.loader').removeClass('loader');
		$('div.loader').remove();
	}, 1000);
}

function getValue(id, validationRegex) {
	var el = $(id);
	if (el == null) return null;
	el.removeClass('invalid');
	if (el.val().match(validationRegex)) return el.val();
	el.addClass('invalid');
	return null;
}

function getValues() {
	var data = {};
	data.name = getValue('#name', /[a-zA-Z0-9 \\.,\\?\\!]+/);
	data.url = getValue('#url', /youtube\.com\/watch\?v=([a-zA-Z0-9\-_]+)|youtu.be\/([a-zA-Z0-9\-]+)/);
	data.start = getValue('#start', /^([0-9]+:[0-9]+:|[0-9]+:|)[0-9]+$/);
	data.duration = getValue('#duration', /[0-9]+/);

	if (data.name == null || data.url == null || data.start == null || data.duration == null) {
		return null;
	}

	data.id = data.url.match(/youtube\.com\/watch\?v=([a-zA-Z0-9\-_]+)|youtu.be\/([a-zA-Z0-9\-]+)/)[1];

	var startParts = data.start.split(':');
	var startSeconds = 0;
	if (startParts.length > 0) startSeconds += parseInt(startParts.pop());
	if (startParts.length > 0) startSeconds += parseInt(startParts.pop()) * 60;
	if (startParts.length > 0) startSeconds += parseInt(startParts.pop()) * 60 * 60;
	data.start = startSeconds;

	data.duration = parseInt(data.duration);

	return data;
}

function prelistenClick() {
	var data = getValues();
	if (data == null) return;
	$('#prel-btn').prepend('<div class="loader"></div>');
	playSample(data.id, data.start, data.duration);
}

function submitClick() {
	var data = getValues();
	if (data == null) return;
	$('#submit-btn').prepend('<div class="loader"></div>');
	$('.close').hide();

	$.post('submit', data).done(function() {
		window.location.reload();
	}).fail(function(){
		alert("Could not submit this entry.");
		$('.close').show();
		$('.loader').remove();
	});
}

function fillList() {
	var userCookie = getCookie('token');
	var user = userCookie ? JSON.parse(atob(userCookie.split('.')[1])) : null;
	$.getJSON('entries', function(data){
		for(x of data) {
			var deletable = user && (user.IsAdmin || (user.OwnerOf && user.OwnerOf.indexOf(x.ID) >= 0));
			var secondLink = `<a target="_blank" href="https://youtu.be/${x.YoutubeID}?t=${x.SecondsStart}s">Open</a>`;
			if (deletable) secondLink = `<a onclick="removeEntry('${x.ID}', this)">Remove</a>`;
			$('#list-body').append(`
				<tr>
					<td>${x.Name}</td>
					<td>${Math.floor(x.SecondsStart/3600).pad(2)}:${Math.floor((x.SecondsStart%3600)/60).pad(2)}:${(x.SecondsStart % 60).pad(2)}</td>
					<td>${x.Duration}s</td>
					<td>
						<a onclick="playSample('${x.YoutubeID}', ${x.SecondsStart}, ${x.Duration}, this)">Listen</a>, 
						${secondLink}
					</td>
				</tr>
			`);
		}
	});
}

function removeEntry(id, sender) {
	if (!confirm('Are you sure you want to delete this thing?')) return;
	if (!confirm('Are you REALLY sure you want to delete this thing???')) return;
	if (sender != null) {
		$(sender).addClass('loader');
	}
	$.post('delete', { 'id': id	}).done(function() {
		$(sender).parent().parent().remove();
		clearLoaders();
	}).fail(function(){
		alert("Could not delete this entry.");
		clearLoaders();
	});
}

function initPage() {
	fillList();
	clearLoaders();
}

$(initPage);