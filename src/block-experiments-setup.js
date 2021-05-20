/**
 * WordPress dependencies
 */
import { select } from '@wordpress/data';

// This file is to set up the jetpack/layout-grid block that currently lives in block-experiments/blocks/layout-grid
import { registerBlock } from '../block-experiments/blocks/layout-grid/src';

export default function setupBlockExperiments() {
	// eslint-disable-next-line @wordpress/react-no-unsafe-timeout
	setTimeout( () => {
		const capabilities = select( 'core/block-editor' ).getSettings(
			'capabilities'
		);
		if ( capabilities.layoutGridBlock ) {
			registerBlock();
		}
	} );
}
