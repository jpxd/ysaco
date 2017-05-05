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

function getValues() {
	$('#name').removeClass('invalid');
	$('#url').removeClass('invalid');
	$('#start').removeClass('invalid');
	$('#duration').removeClass('invalid');

	var data = {};
	data.name = $('#name').val();
	data.url = $('#url').val();
	data.start = $('#start').val();
	data.duration = $('#duration').val();

	if (data.name == null || data.name.length < 5) {
		$('#name').addClass('invalid');
		return null;
	}

	if (data.url == null || data.url.length < 5) {
		$('#url').addClass('invalid');
		return null;
	}

	var urlRes = data.url.match(/youtube\.com\/watch\?v=([a-zA-Z0-9]+)/);
	if (urlRes == null || urlRes.length < 2 || urlRes[1].length < 3) {
		$('#url').addClass('invalid');
		return null;
	}
	data.id = urlRes[1];

	if (data.start == null || data.start.length < 1 || data.start.indexOf(':') < 0) {
		$('#start').addClass('invalid');
		return null;
	}

	var startParts = data.start.split(':');
	if (startParts.length != 2 || isNaN(startParts[0]) || isNaN(startParts[1])) {
		$('#start').addClass('invalid');
		return null;
	}

	data.start = parseInt(startParts[0])*60 + parseInt(startParts[1]);

	if (data.duration == null || data.duration.length < 1 || isNaN(data.duration)) {
		$('#duration').addClass('invalid');
		return null;
	}

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
	}).error(function(){
		alert("Could not submit this entry.");
		$('.close').show();
		$('.loader').remove();
	});
}

function fillList() {
	$.getJSON('entries', function(data){
		for(x of data) {
			$('#list-body').append(`
				<tr>
					<td>${x.Name}</td>
					<td>${x.Duration}</td>
					<td><a onclick="playSample('${x.YoutubeID}', ${x.SecondsStart}, ${x.Duration}, this)">Listen</a></td>
				</tr>
			`);
		}
	});
}

$(function() {
	fillList();
	clearLoaders();
});