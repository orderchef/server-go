var app = angular.module('orderchef');

app.service('reportDates', function () {
	var self = this;
	var dateFormat = 'DD/MM/YYYY';

	this.setup = function (cb) {
		self.onUpdate = cb;
		var pickers = [
			$('.range-start'),
			$('.range-end')
		];

		$('#datepicker').datepicker({
			todayHighlight: true,
			format: 'dd/mm/yyyy',
			inputs: $('.range-start, .range-end')
		});

		pickers[0].datepicker('update', moment().format(dateFormat));
		pickers[1].datepicker('update', moment().format(dateFormat));
		$('#datepicker').datepicker('updateDates');

		pickers[0].on('changeDate', self.startDateChanged);
		pickers[1].on('changeDate', self.endDateChanged);

		var start = localStorage['report_start'];
		var end = localStorage['report_end'];
		var last = localStorage['last_saved'];
		if (!start || !end || !last) return;

		if (last - Date.now() > 86400) {
			start = Date.now()
			end = Date.now()
		}

		self.start = new Date(parseInt(start));
		self.end = new Date(parseInt(end));

		if (self.start.valueOf() == 'NaN' || self.end.valueOf() == 'NaN') {
			return;
		}

		pickers[0].datepicker('update', self.start);
		pickers[1].datepicker('update', self.end);
		$('#datepicker').datepicker('updateDates');
	}

	this.start = new Date;
	this.end = new Date;
	this.onUpdate = null;

	this.startDateChanged = function (ev) {
		self.start = ev.date;
		localStorage['report_start'] = moment(self.start).valueOf();
		localStorage['last_saved'] = Date.now();

		if (typeof self.onUpdate == 'function') self.onUpdate();
	}
	this.endDateChanged = function (ev) {
		self.end = ev.date;
		localStorage['report_end'] = moment(self.end).valueOf();
		localStorage['last_saved'] = Date.now();

		if (typeof self.onUpdate == 'function') self.onUpdate();
	}

	this.getQuery = function () {
		return '?start=' + moment(self.start).unix() + '&end=' + (moment(self.end).unix() + 86400);
	}

	return this;
});