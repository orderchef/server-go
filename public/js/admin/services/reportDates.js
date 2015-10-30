var app = angular.module('orderchef');

app.service('reportDates', function () {
	var self = this;
	var dateFormat = 'DD/MM/YYYY';

	this.setup = function () {
		$('.input-daterange').datepicker({
			orientation: 'top left',
			todayBtn: 'linked',
			autoclose: true,
			format: 'dd/mm/yyyy'
		});

		var pickers = [
			$($('.input-daterange input')[0]),
			$($('.input-daterange input')[1])
		];

		pickers[0].datepicker('update', self.start);
		pickers[1].datepicker('update', self.end);

		pickers[0].on('changeDate', self.datesChanged);
		pickers[1].on('changeDate', self.datesChanged);

		var start = localStorage['report_start'];
		var end = localStorage['report_end'];
		var last = localStorage['last_saved'];
		if (!start || !end || !last) return;

		if (last - Date.now() > 86400) {
			start = Date.now()
			end = Date.now()
		}

		start = new Date(parseInt(start));
		end = new Date(parseInt(end));

		pickers[0].datepicker('update', start);
		pickers[1].datepicker('update', end);

		self.start = moment(start).format(dateFormat);
		self.end = moment(end).format(dateFormat);
	}

	this.start = moment().format(dateFormat);
	this.end = moment().format(dateFormat);

	this.datesChanged = function () {
		localStorage['report_start'] = self.getDate(self.start).valueOf();
		localStorage['report_end'] = self.getDate(self.end).valueOf();
		localStorage['last_saved'] = self.getDate(Date.now());
	}

	this.getDate = function (date) {
		return moment(date, dateFormat);
	}

	this.getQuery = function () {
		return '?start=' + self.getDate(self.start).unix() + '&end=' + (self.getDate(self.end).unix() + 86400);
	}

	return this;
});