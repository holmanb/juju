from mock import patch
import os
from unittest import TestCase

from candidate import (
    find_publish_revision_number,
    get_artifact_dirs,
    prepare_dir,
    update_candidate,
)
from utility import temp_dir


class CandidateTestCase(TestCase):

    @staticmethod
    def make_publish_revision_build_data(*args, **kwargs):
        if kwargs['build'] == 'lastSuccessfulBuild':
            number = 1235
        else:
            number = kwargs['build']
        return {
            'number': number,
            'description': 'Revision build: %s' % number,
        }

    def test_find_publish_revision_number(self):
        with patch('candidate.get_build_data',
                   side_effect=self.make_publish_revision_build_data) as mock:
            found_number = find_publish_revision_number(1234, limit=2)
        self.assertEqual(1234, found_number)
        self.assertEqual(2, mock.call_count)
        output, args, kwargs = mock.mock_calls[0]
        self.assertEqual(
            ('http://juju-ci.vapour.ws:8080', 'publish-revision'), args)
        self.assertEqual('lastSuccessfulBuild', kwargs['build'])
        output, args, kwargs = mock.mock_calls[1]
        self.assertEqual(
            ('http://juju-ci.vapour.ws:8080', 'publish-revision'), args)
        self.assertEqual(1234, kwargs['build'])

    def test_find_publish_revision_number_no_match(self):
        with patch('candidate.get_build_data',
                   side_effect=self.make_publish_revision_build_data) as mock:
            found_number = find_publish_revision_number(1, limit=2)
        self.assertEqual(None, found_number)
        self.assertEqual(2, mock.call_count)

    def test_find_publish_revision_number_no_build_data(self):
        with patch('candidate.get_build_data', return_value=None) as mock:
            found_number = find_publish_revision_number(1, limit=5)
        self.assertEqual(None, found_number)
        self.assertEqual(1, mock.call_count)

    def test_prepare_dir_clean_existing(self):
        with temp_dir() as base_dir:
            candidate_dir_path = os.path.join(base_dir, 'candidate', 'master')
            os.makedirs(candidate_dir_path)
            os.makedirs(os.path.join(candidate_dir_path, 'subdir'))
            candidate_file_path = os.path.join(candidate_dir_path, 'vars.json')
            with open(candidate_file_path, 'w') as cf:
                cf.write('data')
            prepare_dir(candidate_dir_path)
            self.assertTrue(os.path.isdir(candidate_dir_path))
            self.assertEqual([], os.listdir(candidate_dir_path))

    def test_prepare_dir_create_new(self):
        with temp_dir() as base_dir:
            os.makedirs(os.path.join(base_dir, 'candidate'))
            candidate_dir_path = os.path.join(base_dir, 'candidate', 'master')
            prepare_dir(candidate_dir_path)
            self.assertTrue(os.path.isdir(candidate_dir_path))
            self.assertEqual([], os.listdir(candidate_dir_path))

    def test_update_candidate(self):
        with patch('candidate.prepare_dir') as pd_mock:
            with patch('candidate.find_publish_revision_number',
                       return_value=5678) as fprn_mock:
                with patch('candidate.get_artifacts') as ga_mock:
                    update_candidate('gitbr:1.21:gh', '~/candidate', '1234')
        args, kwargs = pd_mock.call_args
        self.assertEqual(('~/candidate/1.21-artifacts', False, False), args)
        args, kwargs = fprn_mock.call_args
        self.assertEqual(('1234', ), args)
        self.assertEqual(2, ga_mock.call_count)
        output, args, kwargs = ga_mock.mock_calls[0]
        self.assertEqual(
            ('build-revision', '1234', 'buildvars.json',
             '~/candidate/1.21-artifacts'),
            args)
        options = {'verbose': False, 'dry_run': False}
        self.assertEqual(options, kwargs)
        output, args, kwargs = ga_mock.mock_calls[1]
        self.assertEqual(
            ('publish-revision', 5678, 'juju-core*',
             '~/candidate/1.21-artifacts'),
            args)
        self.assertEqual(options, kwargs)

    def test_get_artifact_dirs(self):
        with temp_dir() as base_dir:
            os.makedirs(os.path.join(base_dir, 'master'))
            os.makedirs(os.path.join(base_dir, 'master-artifacts'))
            os.makedirs(os.path.join(base_dir, '1.21-artifacts'))
            os.makedirs(os.path.join(base_dir, 'subdir'))
            dirs = get_artifact_dirs(base_dir)
        self.assertEqual(
            ['1.21-artifacts', 'master-artifacts'], sorted(dirs))
