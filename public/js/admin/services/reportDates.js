var app = angular.module('orderchef');

app.service('reportDates', function () {
	var self = this;

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

		pickers[0].datepicker('update', this.start);
		pickers[1].datepicker('update', this.end);

		pickers[0].on('changeDate', self.datesChanged);
		pickers[1].on('changeDate', self.datesChanged);

		var start = localStorage['report_start'];
		var end = localStorage['report_end'];
		if (!start || !end) return;

		start = new Date(parseInt(start));
		end = new Date(parseInt(end));

		pickers[0].datepicker('update', start);
		pickers[1].datepicker('update', end);

		self.start = moment(start).format('DD/MM/YYYY');
		self.end = moment(end).format('DD/MM/YYYY');
	}

	this.start = moment().format('DD/MM/YYYY');
	this.end = moment().format('DD/MM/YYYY');

	this.datesChanged = function () {
		localStorage['report_start'] = moment(self.start, 'DD/MM/YYYY').valueOf();
		localStorage['report_end'] = moment(self.end, 'DD/MM/YYYY').valueOf();
	}

	return this;
});